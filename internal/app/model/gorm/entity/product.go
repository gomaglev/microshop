package entity

import (
	"context"

	"github.com/gomaglev/microshop/pkg/util"

	proto "github.com/gomaglev/protos/pkg/proto/product"
	"gorm.io/gorm"
)

// Product
type Product struct {
	Model
}

// GetProductDB
func GetProductDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, (*gorm.DB)(defDB), new(Product))
}

// TableName
func (e Product) TableName() string {
	return "product"
}

// ToProto
func (e *Product) ToProto() *proto.Product {
	item := new(proto.Product)
	_ = util.StructMapToStruct(e, item)
	return item
}

// Products
type Products []*Product

// ToProtos
func (e Products) ToProtos() []*proto.Product {
	list := make([]*proto.Product, len(e))
	for i, v := range e {
		list[i] = v.ToProto()
	}
	return list
}
