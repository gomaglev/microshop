package model

import (
	"context"

	"github.com/gomaglev/microshop/v1/internal/app/dto"
	"github.com/gomaglev/microshop/v1/internal/app/model"
	"github.com/gomaglev/microshop/v1/internal/app/model/gorm/entity"

	"github.com/gomaglev/protos/pkg/proto/common/v1"
	"github.com/gomaglev/protos/pkg/proto/order/item/message/v1"
	"gorm.io/gorm"

	"github.com/google/wire"
	"github.com/pkg/errors"
)

var _ model.IOrderItemMessage = (*OrderItemMessage)(nil)

// OrderItemMessageSet
var OrderItemMessageSet = wire.NewSet(wire.Struct(new(OrderItemMessage), "*"), wire.Bind(new(model.IOrderItemMessage), new(*OrderItemMessage)))

// OrderItemMessage
type OrderItemMessage struct {
	DB *gorm.DB
}

// List
func (a *OrderItemMessage) List(ctx context.Context, params *dto.ListOrderItemMessagesParam, opts ...*common.QueryOptions) (*message.OrderItemMessages, error) {
	opt := GetQueryOption(opts...)
	db := entity.GetOrderItemMessageDB(ctx, a.DB)
	if v := params.Ids; len(v) > 0 {
		db = db.Where("device_id in (?)", v)
	}

	var list entity.OrderItemMessages
	pr, err := WrapPageQuery(ctx, db, params.Pagination, opt, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &message.OrderItemMessages{
		Pagination: pr,
		List:       list.ToProtos(),
	}
	return qr, nil
}

// Get
func (a *OrderItemMessage) Get(ctx context.Context, params *dto.GetOrderItemMessageParam) (*message.OrderItemMessage, error) {
	var message entity.OrderItemMessage
	ok, err := FindOne(ctx, entity.GetOrderItemMessageDB(ctx, a.DB).Where(`id=?`, params.Id), &message)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return message.ToProto(), nil
}

// Create
func (a *OrderItemMessage) Create(ctx context.Context, item *message.OrderItemMessage) (*common.IDResult, error) {
	pitem := dto.ProtoOrderItemMessage{OrderItemMessage: item}
	eitem := pitem.ToEntity()
	result := entity.GetOrderItemMessageDB(ctx, a.DB).Create(eitem)
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &common.IDResult{
		Id: item.Id,
	}, nil
}

// Update
func (a *OrderItemMessage) Update(ctx context.Context, params *dto.UpdateOrderItemMessageParam, item *message.OrderItemMessage) (*int64, error) {
	pitem := dto.ProtoOrderItemMessage{OrderItemMessage: item}
	eitem := pitem.ToEntity()
	result := entity.GetOrderItemMessageDB(ctx, a.DB).Where("id=?", params.Id).Updates(eitem)
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &result.RowsAffected, nil
}

// UpdateColumns updates order columns
func (a *OrderItemMessage) UpdateColumns(ctx context.Context, params *dto.UpdateOrderItemMessageParam, values map[string]interface{}) (*int64, error) {
	db := entity.GetOrderDB(ctx, a.DB)

	db = db.Where("id = ?", params.Id)

	result := db.UpdateColumns(values)
	if err := result.Error; err != nil {
		return &result.RowsAffected, errors.WithStack(err)
	}
	return &result.RowsAffected, nil
}

// Delete
func (a *OrderItemMessage) Delete(ctx context.Context, params *dto.DeleteOrderItemMessageParam) (*int64, error) {
	result := entity.GetOrderItemMessageDB(ctx, a.DB).Where("id=?", params.Id).Delete(entity.OrderItemMessage{})
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &result.RowsAffected, nil
}
