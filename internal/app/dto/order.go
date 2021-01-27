package dto

import (
	"github.com/gomaglev/microshop/internal/app/model/gorm/entity"
	"github.com/gomaglev/microshop/pkg/util"

	"github.com/gomaglev/protos/pkg/proto/common"
	proto "github.com/gomaglev/protos/pkg/proto/order"
)

// GetOrderParam
type GetOrderParam struct {
	Id string
}

// ListOrdersParam
type ListOrdersParam struct {
	Pagination *common.PaginationParam
	Id         string
	Ids        []string
}

// ListOrdersOptions
type ListOrdersOptions struct {
	OrderByFields []*common.OrderByField
	SelectFields  []string
}

// UpdateOrderParam
type UpdateOrderParam struct {
	Id string
}

// DeleteOrderParam
type DeleteOrderParam struct {
	Id  string
	Ids []string
}

// ProtoOrders
type ProtoOrders proto.Orders

// ToIDs
func (a *ProtoOrders) ToIDs() []string {
	list := make([]string, len(a.List))
	for i, item := range a.List {
		list[i] = item.Id
	}
	return list
}

// ToMap
func (a *ProtoOrders) ToMap() map[string]*proto.Order {
	orderMap := make(map[string]*proto.Order)

	for _, item := range a.List {
		orderMap[item.Id] = item
	}

	return orderMap
}

// ProtoOrder proto type
type ProtoOrder struct {
	Order *proto.Order
}

// ToEntity converts to entity
func (p *ProtoOrder) ToEntity() *entity.Order {
	item := new(entity.Order)
	_ = util.StructMapToStruct(p.Order, item)
	// conversion
	return item
}
