package category

import (
	"context"

	"github.com/gomaglev/microshop/internal/app/dto"
	"github.com/gomaglev/microshop/internal/app/model"
	iutil "github.com/gomaglev/microshop/internal/pkg/util"

	"github.com/google/wire"
)

// make sure Category implements category_grpc CategoryServiceServer
var _ CategoryServiceServer = (*CategoryService)(nil)

// CategorySet injection
var CategorySet = wire.NewSet(wire.Struct(new(CategoryService), "*"), wire.Bind(new(CategoryServiceServer), new(*CategoryService)))

// Category implements CategoryServiceServer
type CategoryService struct {
	CategoryModel model.ICategory
}

// Get category
func (a *CategoryService) Get(ctx context.Context, req *GetCategoryRequest) (*GetCategoryResponse, error) {
	params := &dto.GetCategoryParam{
		Id: req.Id,
	}
	category, err := a.CategoryModel.Get(ctx, params)
	res := &GetCategoryResponse{
		Category: category,
	}
	return res, err
}

// List categories
func (a *CategoryService) List(ctx context.Context, req *ListCategoriesRequest) (*ListCategoriesResponse, error) {
	params := &dto.ListCategoriesParam{
		Ids: req.Ids,
	}
	categories, err := a.CategoryModel.List(ctx, params)
	res := &ListCategoriesResponse{
		Categories: categories,
	}
	return res, err
}

// Create category
func (a *CategoryService) Create(ctx context.Context, req *CreatCategoryRequest) (*CreatCategoryResponse, error) {
	req.Category.Id = iutil.NewID()
	result, err := a.CategoryModel.Create(ctx, req.Category)
	res := &CreatCategoryResponse{
		Id: result.Id,
	}
	return res, err
}

// Update category
func (a *CategoryService) Update(ctx context.Context, req *UpdateCategoryRequest) (*UpdateCategoryResponse, error) {
	params := &dto.UpdateCategoryParam{
		Id: req.Id,
	}
	rows, err := a.CategoryModel.Update(ctx, params, req.Category)
	res := &UpdateCategoryResponse{
		Updated: *rows,
	}
	return res, err
}

// Delete category
func (a *CategoryService) Delete(ctx context.Context, req *DeleteCategoryRequest) (*DeleteCategoryResponse, error) {
	params := &dto.DeleteCategoryParam{
		Id:  req.Id,
		Ids: req.Ids,
	}
	rows, err := a.CategoryModel.Delete(ctx, params)
	res := &DeleteCategoryResponse{
		Deleted: *rows,
	}
	return res, err
}
