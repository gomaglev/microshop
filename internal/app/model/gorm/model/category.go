package model

import (
	"context"

	"github.com/gomaglev/microshop/internal/app/dto"
	"github.com/gomaglev/microshop/internal/app/model"
	"github.com/gomaglev/microshop/internal/app/model/gorm/entity"

	"github.com/gomaglev/microshop/pkg/proto/category"

	"github.com/gomaglev/protos/pkg/proto/common"
	"gorm.io/gorm"

	"github.com/google/wire"
	"github.com/pkg/errors"
)

var _ model.ICategory = (*Category)(nil)

// CategorySet
var CategorySet = wire.NewSet(wire.Struct(new(Category), "*"), wire.Bind(new(model.ICategory), new(*Category)))

// Category
type Category struct {
	DB *gorm.DB
}

// List categories
func (m *Category) List(ctx context.Context, params *dto.ListCategoriesParam, opts ...*common.QueryOptions) (*category.Categories, error) {
	opt := GetQueryOption(opts...)
	db := entity.GetCategoryDB(ctx, m.DB)

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if v := params.Id; v != "" {
		db = db.Where("id = ?", v)
	}

	if v := params.Ids; len(v) > 0 {
		db = db.Where("id in ?", v)
	}

	var list entity.Categories
	pr, err := WrapPageQuery(ctx, db, params.Pagination, opt, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &category.Categories{
		Pagination: pr,
		List:       list.ToProtos(),
	}
	return qr, nil
}

// Get
func (m *Category) Get(ctx context.Context, params *dto.GetCategoryParam) (*category.Category, error) {
	var item entity.Category
	ok, err := FindOne(ctx, entity.GetCategoryDB(ctx, m.DB).Where(`id=?`, params.Id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToProto(), nil
}

// Create
func (m *Category) Create(ctx context.Context, item *category.Category) (*common.IDResult, error) {
	pitem := dto.ProtoCategory{Category: item}
	eitem := pitem.ToEntity()
	result := entity.GetCategoryDB(ctx, m.DB).Create(eitem)
	if err := result.Error; err != nil {

		return nil, errors.WithStack(err)
	}
	return &common.IDResult{
		Id: item.Id,
	}, nil
}

// Update
func (m *Category) Update(ctx context.Context, params *dto.UpdateCategoryParam, item *category.Category) (*int64, error) {
	pitem := dto.ProtoCategory{Category: item}
	eitem := pitem.ToEntity()

	result := entity.GetCategoryDB(ctx, m.DB).Where("id=?", params.Id).Updates(eitem)
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &result.RowsAffected, nil
}

// UpdateColumns updates category columns
func (m *Category) UpdateColumns(ctx context.Context, params *dto.UpdateCategoryParam, values map[string]interface{}) (*int64, error) {
	db := entity.GetCategoryDB(ctx, m.DB)

	db = db.Where("id = ?", params.Id)

	result := db.UpdateColumns(values)
	if err := result.Error; err != nil {
		return &result.RowsAffected, errors.WithStack(err)
	}
	return &result.RowsAffected, nil
}

// Delete
func (m *Category) Delete(ctx context.Context, params *dto.DeleteCategoryParam) (*int64, error) {
	db := entity.GetCategoryDB(ctx, m.DB)
	rowsAffected := int64(0)
	if params.Id != "" {
		db = db.Where("id=?", params.Id)
		result := db.Delete(entity.Category{})
		if err := result.Error; err != nil {
			return nil, errors.WithStack(err)
		}
		rowsAffected = result.RowsAffected
	}

	if len(params.Ids) > 0 {
		db = db.Where("id in (?)", params.Ids)
		result := db.Delete(entity.Category{})
		if err := result.Error; err != nil {
			return nil, errors.WithStack(err)
		}
		rowsAffected = result.RowsAffected
	}

	// add other conditions here
	return &rowsAffected, nil
}
