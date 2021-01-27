package category

import (
	"context"

	"github.com/google/wire"
)

// make sure Category implements category_grpc CategoryServiceServer
var _ CategoryServiceServer = (*Category)(nil)

// CategorySet injection
var CategorySet = wire.NewSet(wire.Struct(new(Category), "*"))

// Category implements CategoryServiceServer
type Category struct {
	//CategoryModel model.ICategory
}

// Get category
func (a *Category) Get(ctx context.Context, req *GetCategoryRequest) (*GetCategoryResponse, error) {
	// params := dto.GetCategoryParam{
	// 	Id: req.Id,
	// }
	// category, err := b.CategoryModel.Get(ctx, params)
	// res := &GetCategoryResponse{
	// 	Category: category,
	// }
	return nil, nil
}

// List categorys
func (a *Category) List(ctx context.Context, req *ListCategoriesRequest) (*ListCategoriesResponse, error) {
	res := &ListCategoriesResponse{
		Categories: nil,
	}
	return res, nil
}

// Create category
func (a *Category) Create(ctx context.Context, req *CreatCategoryRequest) (*CreatCategoryResponse, error) {
	return nil, nil
}

// Update category
func (a *Category) Update(ctx context.Context, req *UpdateCategoryRequest) (*UpdateCategoryResponse, error) {
	return nil, nil
}

// Delete category
func (a *Category) Delete(ctx context.Context, req *DeleteCategoryRequest) (*DeleteCategoryResponse, error) {
	return nil, nil
}
