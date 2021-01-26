package entity

import (
	"context"

	"github.com/gomaglev/microshop/pkg/util"

	proto "github.com/gomaglev/protos/pkg/proto/order/v1"
	"gorm.io/gorm"
)

// Order
type Order struct {
	Model
}

// GetOrderDB
func GetOrderDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, (*gorm.DB)(defDB), new(Order))
}

// TableName
func (e Order) TableName() string {
	return "order"
}

// ToProto
func (e *Order) ToProto() *proto.Order {
	item := new(proto.Order)
	_ = util.StructMapToStruct(e, item)
	return item
}

// Orders
type Orders []*Order

// ToProtos
func (e Orders) ToProtos() []*proto.Order {
	list := make([]*proto.Order, len(e))
	for i, v := range e {
		list[i] = v.ToProto()
	}
	return list
}
