package model

import (
	"context"

	"github.com/gomaglev/microshop/internal/app/dto"

	"github.com/gomaglev/protos/pkg/proto/common/v1"
	"github.com/gomaglev/protos/pkg/proto/order/v1"
)

// IOrder interface for orders
type IOrder interface {
	// List orders
	List(ctx context.Context, params *dto.ListOrdersParam, opts ...*common.QueryOptions) (*order.Orders, error)
	// Get order
	Get(ctx context.Context, params *dto.GetOrderParam) (*order.Order, error)
	// Create order
	Create(ctx context.Context, item *order.Order) (*common.IDResult, error)
	// Update order
	Update(ctx context.Context, params *dto.UpdateOrderParam, item *order.Order) (*int64, error)
	// Update order columns
	UpdateColumns(ctx context.Context, params *dto.UpdateOrderParam, values map[string]interface{}) (*int64, error)
	// Delete order
	Delete(ctx context.Context, params *dto.DeleteOrderParam) (*int64, error)
}
