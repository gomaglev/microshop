package product

import (
	"context"

	"github.com/gomaglev/microshop/internal/app/dto"
	"github.com/gomaglev/microshop/internal/app/model"
	iutil "github.com/gomaglev/microshop/internal/pkg/util"

	"github.com/google/wire"
)

// make sure Product implements product_grpc ProductServiceServer
var _ ProductServiceServer = (*ProductService)(nil)

// ProductSet injection
var ProductSet = wire.NewSet(wire.Struct(new(ProductService), "*"), wire.Bind(new(ProductServiceServer), new(*ProductService)))

// Product implements ProductServiceServer
type ProductService struct {
	ProductModel model.IProduct
}

// Get product
func (a *ProductService) Get(ctx context.Context, req *GetProductRequest) (*GetProductResponse, error) {
	params := &dto.GetProductParam{
		Id: req.Id,
	}
	product, err := a.ProductModel.Get(ctx, params)
	res := &GetProductResponse{
		Product: product,
	}
	return res, err
}

// List products
func (a *ProductService) List(ctx context.Context, req *ListProductsRequest) (*ListProductsResponse, error) {
	params := &dto.ListProductsParam{
		Ids: req.Ids,
	}
	products, err := a.ProductModel.List(ctx, params)
	res := &ListProductsResponse{
		Products: products,
	}
	return res, err
}

// Create product
func (a *ProductService) Create(ctx context.Context, req *CreatProductRequest) (*CreatProductResponse, error) {
	req.Product.Id = iutil.NewID()
	result, err := a.ProductModel.Create(ctx, req.Product)
	res := &CreatProductResponse{
		Id: result.Id,
	}
	return res, err
}

// Update product
func (a *ProductService) Update(ctx context.Context, req *UpdateProductRequest) (*UpdateProductResponse, error) {
	params := &dto.UpdateProductParam{
		Id: req.Id,
	}
	rows, err := a.ProductModel.Update(ctx, params, req.Product)
	res := &UpdateProductResponse{
		Updated: *rows,
	}
	return res, err
}

// Delete product
func (a *ProductService) Delete(ctx context.Context, req *DeleteProductRequest) (*DeleteProductResponse, error) {
	params := &dto.DeleteProductParam{
		Id:  req.Id,
		Ids: req.Ids,
	}
	rows, err := a.ProductModel.Delete(ctx, params)
	res := &DeleteProductResponse{
		Deleted: *rows,
	}
	return res, err
}
