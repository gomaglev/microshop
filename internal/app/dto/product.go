package dto

import (
	"github.com/gomaglev/microshop/internal/app/model/gorm/entity"
	"github.com/gomaglev/microshop/pkg/util"

	"github.com/gomaglev/protos/pkg/proto/common"
	proto "github.com/gomaglev/protos/pkg/proto/product"
)

// GetProductParam
type GetProductParam struct {
	Id string
}

// ListProductsParam
type ListProductsParam struct {
	Pagination *common.PaginationParam
	Id         string
	Ids        []string
}

// ListProductsOptions
type ListProductsOptions struct {
	ProductByFields []*common.OrderByField
	SelectFields    []string
}

// UpdateProductParam
type UpdateProductParam struct {
	Id string
}

// DeleteProductParam
type DeleteProductParam struct {
	Id  string
	Ids []string
}

// ProtoProducts
type ProtoProducts proto.Products

// ToIDs
func (a *ProtoProducts) ToIDs() []string {
	list := make([]string, len(a.List))
	for i, item := range a.List {
		list[i] = item.Id
	}
	return list
}

// ToMap
func (a *ProtoProducts) ToMap() map[string]*proto.Product {
	productMap := make(map[string]*proto.Product)

	for _, item := range a.List {
		productMap[item.Id] = item
	}

	return productMap
}

// ProtoProduct proto type
type ProtoProduct struct {
	Product *proto.Product
}

// ToEntity converts to entity
func (p *ProtoProduct) ToEntity() *entity.Product {
	item := new(entity.Product)
	_ = util.StructMapToStruct(p.Product, item)
	// conversion
	return item
}
