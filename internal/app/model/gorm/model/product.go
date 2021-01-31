package model

import (
	"context"

	"github.com/gomaglev/microshop/internal/app/dto"
	"github.com/gomaglev/microshop/internal/app/model"
	"github.com/gomaglev/microshop/internal/app/model/gorm/entity"

	"github.com/gomaglev/microshop/pkg/proto/common"
	"github.com/gomaglev/microshop/pkg/proto/product"
	"gorm.io/gorm"

	"github.com/google/wire"
	"github.com/pkg/errors"
)

var _ model.IProduct = (*Product)(nil)

// ProductSet
var ProductSet = wire.NewSet(wire.Struct(new(Product), "*"), wire.Bind(new(model.IProduct), new(*Product)))

// Product
type Product struct {
	DB *gorm.DB
}

// List products
func (m *Product) List(ctx context.Context, params *dto.ListProductsParam, opts ...*common.QueryOptions) (*product.Products, error) {
	opt := GetQueryOption(opts...)
	db := entity.GetProductDB(ctx, m.DB)

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if v := params.Id; v != "" {
		db = db.Where("id = ?", v)
	}

	if v := params.Ids; len(v) > 0 {
		db = db.Where("id in ?", v)
	}

	var list entity.Products
	pr, err := WrapPageQuery(ctx, db, params.Pagination, opt, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &product.Products{
		Pagination: pr,
		List:       list.ToProtos(),
	}
	return qr, nil
}

// Get
func (m *Product) Get(ctx context.Context, params *dto.GetProductParam) (*product.Product, error) {
	var item entity.Product
	ok, err := FindOne(ctx, entity.GetProductDB(ctx, m.DB).Where(`id=?`, params.Id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToProto(), nil
}

// Create
func (m *Product) Create(ctx context.Context, item *product.Product) (*common.IDResult, error) {
	pitem := dto.ProtoProduct{Product: item}
	eitem := pitem.ToEntity()
	result := entity.GetProductDB(ctx, m.DB).Create(eitem)
	if err := result.Error; err != nil {

		return nil, errors.WithStack(err)
	}
	return &common.IDResult{
		Id: item.Id,
	}, nil
}

// Update
func (m *Product) Update(ctx context.Context, params *dto.UpdateProductParam, item *product.Product) (*int64, error) {
	pitem := dto.ProtoProduct{Product: item}
	eitem := pitem.ToEntity()

	result := entity.GetProductDB(ctx, m.DB).Where("id=?", params.Id).Updates(eitem)
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &result.RowsAffected, nil
}

// UpdateColumns updates product columns
func (m *Product) UpdateColumns(ctx context.Context, params *dto.UpdateProductParam, values map[string]interface{}) (*int64, error) {
	db := entity.GetProductDB(ctx, m.DB)

	db = db.Where("id = ?", params.Id)

	result := db.UpdateColumns(values)
	if err := result.Error; err != nil {
		return &result.RowsAffected, errors.WithStack(err)
	}
	return &result.RowsAffected, nil
}

// Delete
func (m *Product) Delete(ctx context.Context, params *dto.DeleteProductParam) (*int64, error) {
	db := entity.GetProductDB(ctx, m.DB)
	rowsAffected := int64(0)
	if params.Id != "" {
		db = db.Where("id=?", params.Id)
		result := db.Delete(entity.Product{})
		if err := result.Error; err != nil {
			return nil, errors.WithStack(err)
		}
		rowsAffected = result.RowsAffected
	}

	if len(params.Ids) > 0 {
		db = db.Where("id in (?)", params.Ids)
		result := db.Delete(entity.Product{})
		if err := result.Error; err != nil {
			return nil, errors.WithStack(err)
		}
		rowsAffected = result.RowsAffected
	}

	// add other conditions here
	return &rowsAffected, nil
}
