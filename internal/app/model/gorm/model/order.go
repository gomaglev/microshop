package model

import (
	"context"

	"github.com/gomaglev/microshop/internal/app/dto"
	"github.com/gomaglev/microshop/internal/app/model"
	"github.com/gomaglev/microshop/internal/app/model/gorm/entity"

	"github.com/gomaglev/protos/pkg/proto/common"
	"github.com/gomaglev/protos/pkg/proto/order"
	"gorm.io/gorm"

	"github.com/google/wire"
	"github.com/pkg/errors"
)

var _ model.IOrder = (*Order)(nil)

// OrderSet
var OrderSet = wire.NewSet(wire.Struct(new(Order), "*"), wire.Bind(new(model.IOrder), new(*Order)))

// Order
type Order struct {
	DB *gorm.DB
}

// List orders
func (m *Order) List(ctx context.Context, params *dto.ListOrdersParam, opts ...*common.QueryOptions) (*order.Orders, error) {
	opt := GetQueryOption(opts...)
	db := entity.GetOrderDB(ctx, m.DB)

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if v := params.Id; v != "" {
		db = db.Where("id = ?", v)
	}

	if v := params.Ids; len(v) > 0 {
		db = db.Where("id in ?", v)
	}

	var list entity.Orders
	pr, err := WrapPageQuery(ctx, db, params.Pagination, opt, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &order.Orders{
		Pagination: pr,
		List:       list.ToProtos(),
	}
	return qr, nil
}

// Get
func (m *Order) Get(ctx context.Context, params *dto.GetOrderParam) (*order.Order, error) {
	var item entity.Order
	ok, err := FindOne(ctx, entity.GetOrderDB(ctx, m.DB).Where(`id=?`, params.Id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToProto(), nil
}

// Create
func (m *Order) Create(ctx context.Context, item *order.Order) (*common.IDResult, error) {
	pitem := dto.ProtoOrder{Order: item}
	eitem := pitem.ToEntity()
	result := entity.GetOrderDB(ctx, m.DB).Create(eitem)
	if err := result.Error; err != nil {

		return nil, errors.WithStack(err)
	}
	return &common.IDResult{
		Id: item.Id,
	}, nil
}

// Update
func (m *Order) Update(ctx context.Context, params *dto.UpdateOrderParam, item *order.Order) (*int64, error) {
	pitem := dto.ProtoOrder{Order: item}
	eitem := pitem.ToEntity()

	result := entity.GetOrderDB(ctx, m.DB).Where("id=?", params.Id).Updates(eitem)
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &result.RowsAffected, nil
}

// UpdateColumns updates order columns
func (m *Order) UpdateColumns(ctx context.Context, params *dto.UpdateOrderParam, values map[string]interface{}) (*int64, error) {
	db := entity.GetOrderDB(ctx, m.DB)

	db = db.Where("id = ?", params.Id)

	result := db.UpdateColumns(values)
	if err := result.Error; err != nil {
		return &result.RowsAffected, errors.WithStack(err)
	}
	return &result.RowsAffected, nil
}

// Delete
func (m *Order) Delete(ctx context.Context, params *dto.DeleteOrderParam) (*int64, error) {
	db := entity.GetOrderDB(ctx, m.DB)
	rowsAffected := int64(0)
	if params.Id != "" {
		db = db.Where("id=?", params.Id)
		result := db.Delete(entity.Order{})
		if err := result.Error; err != nil {
			return nil, errors.WithStack(err)
		}
		rowsAffected = result.RowsAffected
	}

	if len(params.Ids) > 0 {
		db = db.Where("id in (?)", params.Ids)
		result := db.Delete(entity.Order{})
		if err := result.Error; err != nil {
			return nil, errors.WithStack(err)
		}
		rowsAffected = result.RowsAffected
	}

	// add other conditions here
	return &rowsAffected, nil
}
