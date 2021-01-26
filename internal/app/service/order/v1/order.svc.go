package order

import (
	"context"

	"github.com/gomaglev/microshop/v1/internal/app/dto"
	"github.com/gomaglev/microshop/v1/internal/app/model"
	iutil "github.com/gomaglev/microshop/v1/internal/pkg/util"

	"github.com/google/wire"
)

// make sure Order implements order_grpc OrderServiceServer
var _ OrderServiceServer = (*OrderService)(nil)

// OrderSet injection
var OrderSet = wire.NewSet(wire.Struct(new(OrderService), "*"), wire.Bind(new(OrderServiceServer), new(*OrderService)))

// Order implements OrderServiceServer
type OrderService struct {
	OrderModel model.IOrder
}

// Get order
func (a *OrderService) Get(ctx context.Context, req *GetOrderRequest) (*GetOrderResponse, error) {
	params := &dto.GetOrderParam{
		Id: req.Id,
	}
	order, err := a.OrderModel.Get(ctx, params)
	res := &GetOrderResponse{
		Order: order,
	}
	return res, err
}

// List orders
func (a *OrderService) List(ctx context.Context, req *ListOrdersRequest) (*ListOrdersResponse, error) {
	params := &dto.ListOrdersParam{
		Ids: req.Ids,
	}
	orders, err := a.OrderModel.List(ctx, params)
	res := &ListOrdersResponse{
		Orders: orders,
	}
	return res, err
}

// Create order
func (a *OrderService) Create(ctx context.Context, req *CreatOrderRequest) (*CreatOrderResponse, error) {
	req.Order.Id = iutil.NewID()
	result, err := a.OrderModel.Create(ctx, req.Order)
	res := &CreatOrderResponse{
		Id: result.Id,
	}
	return res, err
}

// Update order
func (a *OrderService) Update(ctx context.Context, req *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	params := &dto.UpdateOrderParam{
		Id: req.Id,
	}
	rows, err := a.OrderModel.Update(ctx, params, req.Order)
	res := &UpdateOrderResponse{
		Updated: *rows,
	}
	return res, err
}

// Delete order
func (a *OrderService) Delete(ctx context.Context, req *DeleteOrderRequest) (*DeleteOrderResponse, error) {
	params := &dto.DeleteOrderParam{
		Id: req.Id,
	}
	rows, err := a.OrderModel.Delete(ctx, params)
	res := &DeleteOrderResponse{
		Deleted: *rows,
	}
	return res, err
}
