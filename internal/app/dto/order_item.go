package dto

import (
	"github.com/gomaglev/microshop/internal/app/model/gorm/entity"
	"github.com/gomaglev/microshop/pkg/util"

	"github.com/gomaglev/protos/pkg/proto/common"
	proto "github.com/gomaglev/protos/pkg/proto/order/item"
)

// GetOrderItemParam
type GetOrderItemParam struct {
	Id string
}

// ListOrderItemsParam for bll inner usage
type ListOrderItemsParam struct {
	Pagination *common.PaginationParam
	Ids        []string
}

// ListOrderItemsOptions
type ListOrderItemsOptions struct {
	OrderByFields []*common.OrderByField
	SelectFields  []string
}

// UpdateOrderItemParam
type UpdateOrderItemParam struct {
	Id string
}

// DeleteOrderItemParam
type DeleteOrderItemParam struct {
	Id string
}

// ProtoOrderItems
type ProtoOrderItems proto.OrderItems

// ToIDs
func (a *ProtoOrderItems) ToIDs() []string {
	list := make([]string, len(a.List))
	for i, item := range a.List {
		list[i] = item.Id
	}
	return list
}

// ToMap
func (a *ProtoOrderItems) ToMap() map[string]*proto.OrderItem {
	protoMap := make(map[string]*proto.OrderItem)

	for _, item := range a.List {
		protoMap[item.Id] = item
	}

	return protoMap
}

// ProtoOrderItem proto type
type ProtoOrderItem struct {
	OrderItem *proto.OrderItem
}

// ToEntity converts to entity
func (p *ProtoOrderItem) ToEntity() *entity.OrderItem {
	item := new(entity.OrderItem)
	_ = util.StructMapToStruct(p, item)
	// conversion
	return item
}
