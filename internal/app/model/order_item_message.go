package model

import (
	"context"

	"github.com/gomaglev/microshop/internal/app/dto"

	"github.com/gomaglev/protos/pkg/proto/common"
	"github.com/gomaglev/protos/pkg/proto/order/item/message"
)

// IOrderItemMessage interface for device message
type IOrderItemMessage interface {
	// List
	List(ctx context.Context, params *dto.ListOrderItemMessagesParam, opts ...*common.QueryOptions) (*message.OrderItemMessages, error)
	// Get by ID
	Get(ctx context.Context, params *dto.GetOrderItemMessageParam) (*message.OrderItemMessage, error)
	// Create
	Create(ctx context.Context, message *message.OrderItemMessage) (*common.IDResult, error)
	// Update
	Update(ctx context.Context, params *dto.UpdateOrderItemMessageParam, message *message.OrderItemMessage) (*int64, error)
	// Update message columns
	UpdateColumns(ctx context.Context, params *dto.UpdateOrderItemMessageParam, values map[string]interface{}) (*int64, error)
	// Delete
	Delete(ctx context.Context, params *dto.DeleteOrderItemMessageParam) (*int64, error)
}
