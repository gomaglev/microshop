package service

import (
	"context"
	"fmt"

	item "github.com/gomaglev/microshop/internal/app/service/order/item/v1"
	order "github.com/gomaglev/microshop/internal/app/service/order/v1"

	"github.com/gomaglev/microshop/internal/pkg/config"

	"github.com/gomaglev/microshop/internal/pkg/server/rpc"

	"github.com/google/wire"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var _ rpc.IRegister = (*Register)(nil)

// RegisterSet injection
var RegisterSet = wire.NewSet(wire.Struct(new(Register), "*"), wire.Bind(new(rpc.IRegister), new(*Register)))

type Register struct {
	ItemService  *item.ItemService
	OrderService *order.OrderService
}

var (
	targetEndPoint = fmt.Sprintf("localhost:%s", config.C.Gateway.Port)
)

func (r *Register) DialOption() []grpc.DialOption {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	return opts
}

// RegisterServiceServers implementation
func (r *Register) RegisterServiceServers(server *grpc.Server) {
	// Item Service
	item.RegisterItemServiceServer(server, r.ItemService)
	// Order Service
	order.RegisterOrderServiceServer(server, r.OrderService)
}

// RegisterServiceHandlerFromEndpoints implementation
func (r *Register) RegisterServiceHandlerFromEndpoints(ctx context.Context, multiplexer *runtime.ServeMux) error {
	// Item Service
	if err := item.RegisterItemServiceHandlerFromEndpoint(
		ctx, multiplexer, targetEndPoint, r.DialOption()); err != nil {
		return err
	}
	// Order Service
	if err := order.RegisterOrderServiceHandlerFromEndpoint(
		ctx, multiplexer, targetEndPoint, r.DialOption()); err != nil {
		return err
	}
	return nil
}
