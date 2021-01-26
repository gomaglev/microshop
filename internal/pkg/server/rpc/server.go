// Package rpc configures the external gRPC controllers.
package rpc

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/gomaglev/microshop/internal/pkg/config"
	"github.com/gomaglev/microshop/pkg/grpclimit"
	"github.com/gomaglev/microshop/pkg/logger"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"

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

// // Config represents the configuration for gRPC Server.
// type ServerConfig struct {
// 	Host                 string
// 	Port                 string
// 	CertFile             string
// 	KeyFile              string
// 	EnableLogrus         bool
// 	EnableRateLimit      bool
// 	EnableRecovery       bool
// 	EnableAuthentication bool
// 	EnableAuthorization  bool
// 	BasicAuthToken       string

// 	Authentic func(ctx context.Context) (context.Context, error)
// 	Authorize func(ctx context.Context) (context.Context, error)
// }

// Server
type Server struct {
	Register IRegister
	Limiter  grpclimit.Limiter
	//Config   ServerConfig
}

// Setup configures the API package.
func (s *Server) Setup(ctx context.Context) (*grpc.Server, func()) {
	logger.Infof(ctx, "initializing server")
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

	if cfg.Interceptor.EnableRateLimit {
		unaryServerInterceptors = append(unaryServerInterceptors,
			grpc_ratelimit.UnaryServerInterceptor(&s.Limiter))
		streamServerInterceptors = append(streamServerInterceptors,
			grpc_ratelimit.StreamServerInterceptor(&s.Limiter))
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

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(unaryServerInterceptors...),
		grpc_middleware.WithStreamServerChain(streamServerInterceptors...),
	)

	// grpc server registration
	s.Register.RegisterServiceServers(grpcServer)

	bind := fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port)
	lis, err := net.Listen("tcp4", bind)
	if err != nil {
		logger.Fatalf(ctx, "config: unable to listen on port %d, err: %v", cfg.GRPC.Port, err)
	}

	// setup grpc server
	go func() {
		logrus.WithFields(logrus.Fields{"bind": bind}).Info("[API]: starting grpc server")
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatalf(ctx, "failed to serve: %v", err)
		}
	}()

	return grpcServer, func() {
		_, cancel := context.WithTimeout(
			ctx, time.Second*time.Duration(config.C.Gateway.ShutdownTimeout))
		defer cancel()

		grpcServer.GracefulStop()
	}
}
