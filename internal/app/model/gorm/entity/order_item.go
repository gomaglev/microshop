package entity

import (
	"context"

	"github.com/gomaglev/microshop/pkg/util"

	proto "github.com/gomaglev/protos/pkg/proto/order/item/v1"
	"gorm.io/gorm"
)

type OrderItem struct {
	Model
	OrderId string
}

// GetOrderItemDB
func GetOrderItemDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, (*gorm.DB)(defDB), new(OrderItem))
}

// TableName
func (e OrderItem) TableName() string {
	return "order_item"
}

// ToProto
func (e *OrderItem) ToProto() *proto.OrderItem {
	item := new(proto.OrderItem)
	_ = util.StructMapToStruct(e, item)

	return item
}

// OrderItems
type OrderItems []*OrderItem

// ToProtos
func (e OrderItems) ToProtos() []*proto.OrderItem {
	list := make([]*proto.OrderItem, len(e))
	for i, v := range e {
		list[i] = v.ToProto()
	}
	return list
}
