package order

import (
	"context"

	"github.com/gomaglev/microshop/internal/app/dto"
	"github.com/gomaglev/microshop/internal/app/model"

	"github.com/google/wire"
)

// make sure Order implements order_grpc OrderServiceServer
var _ OrderServiceServer = (*OrderService)(nil)

// OrderSet injection
var OrderSet = wire.NewSet(wire.Struct(new(OrderService), "*"))

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
	res := &ListOrdersResponse{
		Orders: nil,
	}
	return res, nil
}

// Create order
func (a *OrderService) Create(ctx context.Context, req *CreatOrderRequest) (*CreatOrderResponse, error) {
	return nil, nil
}

// Update order
func (a *OrderService) Update(ctx context.Context, req *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	return nil, nil
}

// Delete order
func (a *OrderService) Delete(ctx context.Context, req *DeleteOrderRequest) (*DeleteOrderResponse, error) {
	return nil, nil
}
