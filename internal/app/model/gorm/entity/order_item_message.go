package entity

import (
	"context"

	"github.com/gomaglev/microshop/pkg/util"

	proto "github.com/gomaglev/protos/pkg/proto/order/item/message/v1"
	"gorm.io/gorm"
)

// OrderItemMessage
type OrderItemMessage struct {
	OrderId   string
	ItemId    string
	MessageId string
}

// GetOrderItemMessageDB
func GetOrderItemMessageDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, (*gorm.DB)(defDB), new(OrderItemMessage))
}

// OrderItemMessageProto
type OrderItemMessageProto proto.OrderItemMessage

// TableName
func (e OrderItemMessage) TableName() string {
	return "order_item_message"
}

// ToProto
func (e *OrderItemMessage) ToProto() *proto.OrderItemMessage {
	item := new(proto.OrderItemMessage)
	_ = util.StructMapToStruct(e, item)

	return item
}

// OrderItemMessages
type OrderItemMessages []*OrderItemMessage

// ToProtos
func (e OrderItemMessages) ToProtos() []*proto.OrderItemMessage {
	list := make([]*proto.OrderItemMessage, len(e))
	for i, v := range e {
		list[i] = v.ToProto()
	}
	return list
}
