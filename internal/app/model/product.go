package model

import (
	"context"

	"github.com/gomaglev/microshop/internal/app/dto"

	"github.com/gomaglev/microshop/pkg/proto/common"
	"github.com/gomaglev/microshop/pkg/proto/product"
)

// IProduct interface for Product model
type IProduct interface {
	// List   products
	List(ctx context.Context, params *dto.ListProductsParam, opts ...*common.QueryOptions) (*product.Products, error)
	// Get    product
	Get(ctx context.Context, params *dto.GetProductParam) (*product.Product, error)
	// Create product
	Create(ctx context.Context, item *product.Product) (*common.IDResult, error)
	// Update product
	Update(ctx context.Context, params *dto.UpdateProductParam, item *product.Product) (*int64, error)
	// Update product by columns
	UpdateColumns(ctx context.Context, params *dto.UpdateProductParam, values map[string]interface{}) (*int64, error)
	// Delete product
	Delete(ctx context.Context, params *dto.DeleteProductParam) (*int64, error)
}
