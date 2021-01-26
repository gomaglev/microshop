package rpc

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// IRegister interface for api services registration
type IRegister interface {
	// RegisterServiceServers
	RegisterServiceServers(server *grpc.Server)

	// RegisterServiceHandlerFromEndpoints
	RegisterServiceHandlerFromEndpoints(ctx context.Context, multiplexer *runtime.ServeMux) error
}
