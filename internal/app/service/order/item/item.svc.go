package item

import (
	"context"

	"github.com/gomaglev/microshop/internal/app/dto"
	"github.com/gomaglev/microshop/internal/app/model"

	"github.com/google/wire"
)

// make sure Item implements order_grpc ItemServiceServer
var _ ItemServiceServer = (*ItemService)(nil)

// ItemSet injection
var ItemSet = wire.NewSet(wire.Struct(new(ItemService), "*"))

// Item implements ItemServiceServer
type ItemService struct {
	ItemModel model.IOrderItem
}

// Get item
func (a *ItemService) Get(ctx context.Context, req *GetItemRequest) (*GetItemResponse, error) {
	params := &dto.GetOrderItemParam{
		Id: req.Id,
	}
	data, err := a.ItemModel.Get(ctx, params)
	res := &GetItemResponse{
		Item: data,
	}
	return res, err
}

// List items
func (a *ItemService) List(ctx context.Context, req *ListItemsRequest) (*ListItemsResponse, error) {
	res := &ListItemsResponse{
		Items: nil,
	}
	return res, nil
}

// Create item
func (a *ItemService) Create(ctx context.Context, req *CreatItemRequest) (*CreatItemResponse, error) {
	return nil, nil
}

// Update item
func (a *ItemService) Update(ctx context.Context, req *UpdateItemRequest) (*UpdateItemResponse, error) {
	return nil, nil
}

// Delete item
func (a *ItemService) Delete(ctx context.Context, req *DeleteItemRequest) (*DeleteItemResponse, error) {
	return nil, nil
}
