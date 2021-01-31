package model

import (
	"context"

	"github.com/gomaglev/microshop/internal/app/dto"

	"github.com/gomaglev/microshop/pkg/proto/category"

	"github.com/gomaglev/protos/pkg/proto/common"
)

// ICategory interface for Category model
type ICategory interface {
	// List   categories
	List(ctx context.Context, params *dto.ListCategoriesParam, opts ...*common.QueryOptions) (*category.Categories, error)
	// Get    category
	Get(ctx context.Context, params *dto.GetCategoryParam) (*category.Category, error)
	// Create category
	Create(ctx context.Context, item *category.Category) (*common.IDResult, error)
	// Update category
	Update(ctx context.Context, params *dto.UpdateCategoryParam, item *category.Category) (*int64, error)
	// Update category by columns
	UpdateColumns(ctx context.Context, params *dto.UpdateCategoryParam, values map[string]interface{}) (*int64, error)
	// Delete category
	Delete(ctx context.Context, params *dto.DeleteCategoryParam) (*int64, error)
}
