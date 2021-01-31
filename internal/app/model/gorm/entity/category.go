package entity

import (
	"context"

	"github.com/gomaglev/microshop/pkg/util"

	proto "github.com/gomaglev/microshop/pkg/proto/category"

	"gorm.io/gorm"
)

// Category
type Category struct {
	Model
}

// GetCategoryDB
func GetCategoryDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, (*gorm.DB)(defDB), new(Category))
}

// TableName
func (e Category) TableName() string {
	return "category"
}

// ToProto
func (e *Category) ToProto() *proto.Category {
	item := new(proto.Category)
	_ = util.StructMapToStruct(e, item)
	return item
}

// Categories
type Categories []*Category

// ToProtos
func (e Categories) ToProtos() []*proto.Category {
	list := make([]*proto.Category, len(e))
	for i, v := range e {
		list[i] = v.ToProto()
	}
	return list
}
