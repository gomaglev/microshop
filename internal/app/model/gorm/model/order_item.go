package model

import (
	"context"

	"github.com/gomaglev/microshop/internal/app/dto"
	"github.com/gomaglev/microshop/internal/app/model"
	"github.com/gomaglev/microshop/internal/app/model/gorm/entity"

	"github.com/gomaglev/protos/pkg/proto/common"
	"github.com/gomaglev/protos/pkg/proto/order/item"
	"gorm.io/gorm"

	"github.com/google/wire"
	"github.com/pkg/errors"
)

var _ model.IOrderItem = (*OrderItem)(nil)

// OrderItemSet
var OrderItemSet = wire.NewSet(wire.Struct(new(OrderItem), "*"), wire.Bind(new(model.IOrderItem), new(*OrderItem)))

// OrderItem
type OrderItem struct {
	DB *gorm.DB
}

// List
func (a *OrderItem) List(ctx context.Context, params *dto.ListOrderItemsParam, opts ...*common.QueryOptions) (*item.OrderItems, error) {
	opt := GetQueryOption(opts...)
	db := entity.GetOrderItemDB(ctx, a.DB)
	if v := params.Ids; len(v) > 0 {
		db = db.Where("device_id in (?)", v)
	}

	var list entity.OrderItems
	pr, err := WrapPageQuery(ctx, db, params.Pagination, opt, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &item.OrderItems{
		Pagination: pr,
		List:       list.ToProtos(),
	}
	return qr, nil
}

// Get
func (a *OrderItem) Get(ctx context.Context, params *dto.GetOrderItemParam) (*item.OrderItem, error) {
	var item entity.OrderItem
	ok, err := FindOne(ctx, entity.GetOrderItemDB(ctx, a.DB).Where(`id=?`, params.Id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToProto(), nil
}

// Create
func (a *OrderItem) Create(ctx context.Context, item *item.OrderItem) (*common.IDResult, error) {
	pitem := dto.ProtoOrderItem{OrderItem: item}
	eitem := pitem.ToEntity()
	result := entity.GetOrderItemDB(ctx, a.DB).Create(eitem)
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &common.IDResult{
		Id: item.Id,
	}, nil
}

// Update
func (a *OrderItem) Update(ctx context.Context, params *dto.UpdateOrderItemParam, item *item.OrderItem) (*int64, error) {
	pitem := dto.ProtoOrderItem{OrderItem: item}
	eitem := pitem.ToEntity()
	result := entity.GetOrderItemDB(ctx, a.DB).Where("id=?", params.Id).Updates(eitem)
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &result.RowsAffected, nil
}

// UpdateColumns updates order columns
func (m *OrderItem) UpdateColumns(ctx context.Context, params *dto.UpdateOrderItemParam, values map[string]interface{}) (*int64, error) {
	db := entity.GetOrderDB(ctx, m.DB)

	db = db.Where("id = ?", params.Id)

	result := db.UpdateColumns(values)
	if err := result.Error; err != nil {
		return &result.RowsAffected, errors.WithStack(err)
	}
	return &result.RowsAffected, nil
}

// Delete
func (a *OrderItem) Delete(ctx context.Context, params *dto.DeleteOrderItemParam) (*int64, error) {
	result := entity.GetOrderItemDB(ctx, a.DB).Where("id=?", params.Id).Delete(entity.OrderItem{})
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &result.RowsAffected, nil
}
