// Package rpc configures the external gRPC controllers.
package rpc

import (
	"context"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gomaglev/microshop/internal/pkg/config"
	"github.com/gomaglev/microshop/pkg/grpclimit"
	"github.com/gomaglev/microshop/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"go.uber.org/ratelimit"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/google/wire"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ratelimit "github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ServerSet injection
var ServerSet = wire.NewSet(wire.Struct(new(Server), "*"))

// Server
type Server struct {
	Register IRegister
}

// Setup configures grpc, gateway server and register services
func (s *Server) Setup(ctx context.Context) func() {
	logger.Infof(ctx, "initializing server")
	// new grpc server
	grpcServer := s.newGrpcServer()

	// service registration
	s.Register.RegisterServiceServers(grpcServer)

	go func() {
		bind := fmt.Sprintf("%s:%d", config.C.GRPC.Host, config.C.GRPC.Port)
		listener, err := net.Listen("tcp4", bind)
		if err != nil {
			panic(errors.Wrap(err, fmt.Sprintf("config: unable to listen on %s", bind)))
		}

		logrus.WithFields(logrus.Fields{
			"grpc_bind": bind,
		}).Info("starting grpc server")

		logrus.Fatal(grpcServer.Serve(listener))
	}()

	if !config.C.Gateway.Enable {
		return func() {
			grpcServer.GracefulStop()
		}
	}
	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{},
	))

	err := s.Register.RegisterServiceHandlerFromEndpoints(context.Background(), gwmux)
	if err != nil {
		logrus.Fatalf("failed to register handlers: %v", err)
	}

	// http server
	router := mux.NewRouter()
	router.PathPrefix(
		fmt.Sprintf("/{version:v[0-9]+}%s", config.C.Gateway.PathPrefix),
	).Handler(gwmux)

	router.PathPrefix(
		fmt.Sprintf("/{version:v[0-9]+}%s/.*swagger.json", config.C.Gateway.PathPrefix),
	).Handler(s.swaggerHandlerFunc(ctx))

	bind := fmt.Sprintf("%s:%v", config.C.Gateway.Host, config.C.Gateway.Port)
	listener, err := net.Listen("tcp4", bind)
	if err != nil {
		panic(errors.Wrap(err, fmt.Sprintf("config: unable to listen on %s", bind)))
	}

	// grpc gateway server
	srv := &http.Server{
		Addr:    bind,
		Handler: s.grpcHandlerFunc(grpcServer, wsproxy.WebsocketProxy(router)),
	}

	go func() {
		logrus.WithFields(logrus.Fields{
			"gtwy_bind": bind,
		}).Info("starting gateway server")
		logrus.Fatal(srv.Serve(listener))
	}()

	return func() {
		ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
		}()
		if err := srv.Shutdown(ctxShutDown); err != nil {
			logrus.Fatalf("Server Shutdown Failed:%+v", err)
		}
		grpcServer.GracefulStop()
	}
}

func (s *Server) newGrpcServer() *grpc.Server {
	var unaryServerInterceptors = []grpc.UnaryServerInterceptor{}
	var streamServerInterceptors = []grpc.StreamServerInterceptor{}
	cfg := config.C

	if cfg.Interceptor.EnableLogrus {
		logEntry := logrus.NewEntry(logrus.StandardLogger())
		logOpts := []grpc_logrus.Option{
			grpc_logrus.WithLevels(grpc_logrus.DefaultCodeToLevel),
			grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
				return "grpc.time_ms", duration.Milliseconds()
			}),
		}

		unaryServerInterceptors = append(unaryServerInterceptors,
			grpc_logrus.UnaryServerInterceptor(logEntry, logOpts...))
		streamServerInterceptors = append(streamServerInterceptors,
			grpc_logrus.StreamServerInterceptor(logEntry, logOpts...))
	}

	limiter := grpclimit.Limiter{
		Limiter: ratelimit.New(config.C.GRPC.RateLimitCount),
	}

	if cfg.Interceptor.EnableRateLimit {
		unaryServerInterceptors = append(unaryServerInterceptors,
			grpc_ratelimit.UnaryServerInterceptor(&limiter))
		streamServerInterceptors = append(streamServerInterceptors,
			grpc_ratelimit.StreamServerInterceptor(&limiter))
	}

	if cfg.Interceptor.EnableRecovery {
		unaryServerInterceptors = append(unaryServerInterceptors,
			grpc_recovery.UnaryServerInterceptor(
				grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
					return status.Errorf(codes.Unknown, "panic triggered: %v", p)
				}),
			),
		)

		streamServerInterceptors = append(streamServerInterceptors,
			grpc_recovery.StreamServerInterceptor(
				grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
					return status.Errorf(codes.Unknown, "panic triggered: %v", p)
				}),
			),
		)
	}

	if cfg.Interceptor.EnableAuthentication {
		cfg.BasicAuth.AuthToken = base64.StdEncoding.EncodeToString(
			[]byte(fmt.Sprintf("%s:%s", cfg.BasicAuth.User, cfg.BasicAuth.Password)))

		unaryServerInterceptors = append(unaryServerInterceptors,
			grpc_auth.UnaryServerInterceptor(cfg.BasicAuth.AuthFunc()))
		streamServerInterceptors = append(streamServerInterceptors,
			grpc_auth.StreamServerInterceptor(cfg.BasicAuth.AuthFunc()))
	}

	if cfg.Interceptor.EnableAuthorization {
		unaryServerInterceptors = append(unaryServerInterceptors,
			grpc_auth.UnaryServerInterceptor(cfg.Authorizer.Enforce()))
		streamServerInterceptors = append(streamServerInterceptors,
			grpc_auth.StreamServerInterceptor(cfg.Authorizer.Enforce()))
	}

	return grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(unaryServerInterceptors...),
		grpc_middleware.WithStreamServerChain(streamServerInterceptors...),
	)
}

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied and changed from cockroachdb.
func (s *Server) grpcHandlerFunc(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		protoMajor := r.ProtoMajor
		urlPath := r.URL.Path
		logrus.Printf("%s,%v,%s", contentType, protoMajor, urlPath)
		if protoMajor == 2 && strings.Contains(contentType, "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			if httpHandler == nil {
				w.WriteHeader(http.StatusNotImplemented)
				return
			}
			if urlPath == "/status" {
				w.WriteHeader(http.StatusOK)
				return
			}

			s.corsHandleFunc(httpHandler).ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func (s *Server) swaggerHandlerFunc(ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
			logger.Errorf(ctx, "Swagger JSON not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		logger.Infof(ctx, "Serving %s", r.URL.Path)
		p := strings.TrimPrefix(r.URL.Path, "SITE URL")
		http.ServeFile(w, r, p)
	})
}

// corsHandleFunc allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func (s *Server) corsHandleFunc(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				headers := []string{"Content-Type", "Accept"}
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
				methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}
