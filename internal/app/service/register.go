package service

import (
	"context"
	"fmt"

	orderv1 "github.com/gomaglev/microshop/internal/app/service/v1/order"
	itemv1 "github.com/gomaglev/microshop/internal/app/service/v1/order/item"
	orderv2 "github.com/gomaglev/microshop/internal/app/service/v2/order"
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
	ItemServiceV1  *itemv1.ItemService
	OrderServiceV1 *orderv1.OrderService
	OrderServiceV2 *orderv2.OrderService
}

func (r *Register) DialOption() []grpc.DialOption {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	return opts
}

// RegisterServiceServers implementation
func (r *Register) RegisterServiceServers(server *grpc.Server) {
	// Item Service
	itemv1.RegisterItemServiceServer(server, r.ItemServiceV1)
	// Order Service
	orderv1.RegisterOrderServiceServer(server, r.OrderServiceV1)
	// Order Service
	orderv2.RegisterOrderServiceServer(server, r.OrderServiceV2)
}

// RegisterServiceHandlerFromEndpoints implementation
func (r *Register) RegisterServiceHandlerFromEndpoints(ctx context.Context, multiplexer *runtime.ServeMux) error {
	endpoint := r.endpoint()
	dialOption := r.DialOption()
	// Item Service
	if err := itemv1.RegisterItemServiceHandlerFromEndpoint(
		ctx, multiplexer, endpoint, dialOption); err != nil {
		return err
	}
	// Order Service
	if err := orderv1.RegisterOrderServiceHandlerFromEndpoint(
		ctx, multiplexer, endpoint, dialOption); err != nil {
		return err
	}
	if err := orderv2.RegisterOrderServiceHandlerFromEndpoint(
		ctx, multiplexer, endpoint, dialOption); err != nil {
		return err
	}
	return nil
}

func (r *Register) endpoint() string {
	return fmt.Sprintf("%s:%d", config.C.Gateway.Host, config.C.Gateway.Port)
}
