package model

import (
	"context"

	"github.com/gomaglev/microshop/v1/internal/app/dto"

	"github.com/gomaglev/protos/pkg/proto/common/v1"
	"github.com/gomaglev/protos/pkg/proto/order/item/v1"
)

// IOrderItem interface for order_items
type IOrderItem interface {
	// Query
	List(ctx context.Context, params *dto.ListOrderItemsParam, opts ...*common.QueryOptions) (*item.OrderItems, error)
	// Get by ID
	Get(ctx context.Context, params *dto.GetOrderItemParam) (*item.OrderItem, error)
	// Create
	Create(ctx context.Context, item *item.OrderItem) (*common.IDResult, error)
	// Update
	Update(ctx context.Context, params *dto.UpdateOrderItemParam, item *item.OrderItem) (*int64, error)
	// Update columns
	UpdateColumns(ctx context.Context, params *dto.UpdateOrderItemParam, values map[string]interface{}) (*int64, error)
	// Delete
	Delete(ctx context.Context, params *dto.DeleteOrderItemParam) (*int64, error)
}
