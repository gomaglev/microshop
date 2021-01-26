// Package rpc configures the external gRPC controllers.
package rpc

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gomaglev/microshop/v1/internal/pkg/config"
	"github.com/gomaglev/microshop/v1/pkg/logger"

	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

// GatewaySet injection
var GatewaySet = wire.NewSet(wire.Struct(new(Gateway), "*"))

type Gateway struct {
	Register IRegister
}

func (s *Gateway) Setup(
	ctx context.Context,
	grpcServer *grpc.Server) func() {

	cfg := config.C

	addr := fmt.Sprintf("%s:%d", cfg.Gateway.Host, cfg.Gateway.Port)

	serveMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{},
		),
	)

	/*
		dcreds := credentials.NewTLS(&tls.Config{
			ServerName: addr,
			RootCAs:    demoCertPool,
		})
		dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	*/

	err := s.Register.RegisterServiceHandlerFromEndpoints(context.Background(), serveMux)
	if err != nil {
		logger.Fatalf(ctx, "failed to register handlers: %v", err)
	}

	router := mux.NewRouter()
	router.PathPrefix(
		fmt.Sprintf("/{version:v[0-9]+}%s", cfg.Gateway.PathPrefix),
	).Handler(serveMux)

	router.PathPrefix(
		fmt.Sprintf("/{version:v[0-9]+}%s/.*swagger.json", cfg.Gateway.PathPrefix),
	).Handler(swaggerHandlerFunc(ctx))

	srv := &http.Server{
		Addr:    addr,
		Handler: s.grpcHandlerFunc(grpcServer, router),
		// TLSConfig: &tls.Config{
		// 	Certificates: []tls.Certificate{
		// 		tls.Certificate{},
		// 	},
		// 	NextProtos:   []string{"h2"},
		// },
	}

	go func() {
		// GRPC_GO_LOG_VERBOSITY_LEVEL
		logger.Infof(ctx, "[API]: starting api server")
		err := srv.ListenAndServe()
		if err != nil {
			logger.Errorf(ctx, err.Error())
		}
	}()

	return func() {
		ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
		}()
		if err := srv.Shutdown(ctxShutDown); err != nil {
			logger.Fatalf(ctx, "Server Shutdown Failed:%+v", err)
		}
	}
}

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied and changed from cockroachdb grpc_util.go.
// https://github.com/cockroachdb/cockroach/blob/master/LICENSE
func (s *Gateway) grpcHandlerFunc(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		protoMajor := r.ProtoMajor
		urlPath := r.URL.Path

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

			corsHandleFunc(httpHandler).ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

// corsHandleFunc allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func corsHandleFunc(h http.Handler) http.Handler {
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

func swaggerHandlerFunc(ctx context.Context) http.Handler {
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

// func (s *Gateway) Setup(
// 	ctx context.Context,
// 	grpcServer *grpc.Server) func() {
// 	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(
// 		runtime.MIMEWildcard,
// 		&runtime.JSONPb{},
// 	))

// 	err := s.Register.RegisterServiceHandlerFromEndpoints(context.Background(), gwmux)
// 	if err != nil {
// 		logger.Fatalf(ctx, "failed to register handlers: %v", err)
// 	}

// 	// http server
// 	router := mux.NewRouter()
// 	router.PathPrefix(fmt.Sprintf("/{version:v[0-9]+}%s", s.Config.PathPrefix)).Handler(gwmux)

// 	srv := &http.Server{
// 		Addr:    s.Config.Host + ":" + s.Config.Port,
// 		Handler: s.grpcHandlerFunc(grpcServer, router, s.Config),
// 	}

// 	go (func() error {
// 		logger.Infof(ctx, "[API]: starting api server")
// 		return srv.ListenAndServe()
// 	})()

// 	return func() {
// 		ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 		defer func() {
// 			cancel()
// 		}()
// 		if err := srv.Shutdown(ctxShutDown); err != nil {
// 			logger.Fatalf(ctx, "Server Shutdown Failed:%+v", err)
// 		}
// 	}
// }

// // grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// // connections or otherHandler otherwise. Copied and changed from cockroachdb.
// func (s *Gateway) grpcHandlerFunc(grpcServer *grpc.Server, httpHandler http.Handler, config GatewayConfig) http.Handler {
// 	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		contentType := r.Header.Get("Content-Type")
// 		protoMajor := r.ProtoMajor
// 		urlPath := r.URL.Path
// 		if protoMajor == 2 && strings.Contains(contentType, "application/grpc") {
// 			grpcServer.ServeHTTP(w, r)
// 		} else {
// 			if httpHandler == nil {
// 				w.WriteHeader(http.StatusNotImplemented)
// 				return
// 			}
// 			if urlPath == "/status" {
// 				w.WriteHeader(http.StatusOK)
// 				return
// 			}

// 			httpHandler.ServeHTTP(w, r)
// 		}
// 	}), &http2.Server{})
// }
