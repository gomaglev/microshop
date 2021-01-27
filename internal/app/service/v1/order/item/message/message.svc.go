package message

import (
	"context"

	"github.com/gomaglev/microshop/internal/app/dto"
	"github.com/gomaglev/microshop/internal/app/model"

	"github.com/google/wire"
)

// make sure Message implements order_grpc MessageServiceServer
var _ MessageServiceServer = (*MessageService)(nil)

// MessageSet injection
var MessageSet = wire.NewSet(wire.Struct(new(MessageService), "*"))

// Message implements MessageServiceServer
type MessageService struct {
	MessageModel model.IOrderItemMessage
}

// Get message
func (a *MessageService) Get(ctx context.Context, req *GetMessageRequest) (*GetMessageResponse, error) {
	params := &dto.GetOrderItemMessageParam{
		Id: req.Id,
	}
	data, err := a.MessageModel.Get(ctx, params)
	res := &GetMessageResponse{
		Message: data,
	}
	return res, err
}

// List messages
func (a *MessageService) List(ctx context.Context, req *ListMessagesRequest) (*ListMessagesResponse, error) {
	res := &ListMessagesResponse{
		Messages: nil,
	}
	return res, nil
}

// Create message
func (a *MessageService) Create(ctx context.Context, req *CreatMessageRequest) (*CreatMessageResponse, error) {
	return nil, nil
}

// Update message
func (a *MessageService) Update(ctx context.Context, req *UpdateMessageRequest) (*UpdateMessageResponse, error) {
	return nil, nil
}

// Delete message
func (a *MessageService) Delete(ctx context.Context, req *DeleteMessageRequest) (*DeleteMessageResponse, error) {
	return nil, nil
}
