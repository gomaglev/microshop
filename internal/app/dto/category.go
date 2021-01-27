package dto

import (
	"github.com/gomaglev/microshop/internal/app/model/gorm/entity"
	"github.com/gomaglev/microshop/pkg/util"

	proto "github.com/gomaglev/protos/pkg/proto/category"
	"github.com/gomaglev/protos/pkg/proto/common"
)

// GetCategoryParam
type GetCategoryParam struct {
	Id string
}

// ListCategoriesParam
type ListCategoriesParam struct {
	Pagination *common.PaginationParam
	Id         string
	Ids        []string
}

// ListCategoriesOptions
type ListCategoriesOptions struct {
	CategoryByFields []*common.OrderByField
	SelectFields     []string
}

// UpdateCategoryParam
type UpdateCategoryParam struct {
	Id string
}

// DeleteCategoryParam
type DeleteCategoryParam struct {
	Id  string
	Ids []string
}

// ProtoCategories
type ProtoCategories proto.Categories

// ToIDs
func (a *ProtoCategories) ToIDs() []string {
	list := make([]string, len(a.List))
	for i, item := range a.List {
		list[i] = item.Id
	}
	return list
}

// ToMap
func (a *ProtoCategories) ToMap() map[string]*proto.Category {
	categoryMap := make(map[string]*proto.Category)

	for _, item := range a.List {
		categoryMap[item.Id] = item
	}

	return categoryMap
}

// ProtoCategory proto type
type ProtoCategory struct {
	Category *proto.Category
}

// ToEntity converts to entity
func (p *ProtoCategory) ToEntity() *entity.Category {
	item := new(entity.Category)
	_ = util.StructMapToStruct(p.Category, item)
	// conversion
	return item
}
