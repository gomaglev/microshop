package model

import (
	"context"

	"github.com/gomaglev/microshop/internal/app/dto"

	"github.com/gomaglev/protos/pkg/proto/common"
	"github.com/gomaglev/protos/pkg/proto/order/item/message"
)

// IOrderItemMessage interface for OrderItemMessage model
type IOrderItemMessage interface {
	// List   message
	List(ctx context.Context, params *dto.ListOrderItemMessagesParam, opts ...*common.QueryOptions) (*message.OrderItemMessages, error)
	// Get    message
	Get(ctx context.Context, params *dto.GetOrderItemMessageParam) (*message.OrderItemMessage, error)
	// Create message
	Create(ctx context.Context, message *message.OrderItemMessage) (*common.IDResult, error)
	// Update message
	Update(ctx context.Context, params *dto.UpdateOrderItemMessageParam, message *message.OrderItemMessage) (*int64, error)
	// Update message by columns
	UpdateColumns(ctx context.Context, params *dto.UpdateOrderItemMessageParam, values map[string]interface{}) (*int64, error)
	// Delete message
	Delete(ctx context.Context, params *dto.DeleteOrderItemMessageParam) (*int64, error)
}
