// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: api/order/item/item.proto

// @protomicro: module:"order/item" plural="orders/items" description:"OrderItem service"

package item

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	common "github.com/gomaglev/protos/pkg/proto/common"
	item "github.com/gomaglev/protos/pkg/proto/order/item"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId string `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *GetItemRequest) Reset() {
	*x = GetItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_item_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemRequest) ProtoMessage() {}

func (x *GetItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_item_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemRequest.ProtoReflect.Descriptor instead.
func (*GetItemRequest) Descriptor() ([]byte, []int) {
	return file_api_order_item_item_proto_rawDescGZIP(), []int{0}
}

func (x *GetItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetItemRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type GetItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *item.OrderItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *GetItemResponse) Reset() {
	*x = GetItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_item_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemResponse) ProtoMessage() {}

func (x *GetItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_item_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemResponse.ProtoReflect.Descriptor instead.
func (*GetItemResponse) Descriptor() ([]byte, []int) {
	return file_api_order_item_item_proto_rawDescGZIP(), []int{1}
}

func (x *GetItemResponse) GetItem() *item.OrderItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type ListItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids        []string                `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	OrderId    string                  `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Pagination *common.PaginationParam `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListItemsRequest) Reset() {
	*x = ListItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_item_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsRequest) ProtoMessage() {}

func (x *ListItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_item_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsRequest.ProtoReflect.Descriptor instead.
func (*ListItemsRequest) Descriptor() ([]byte, []int) {
	return file_api_order_item_item_proto_rawDescGZIP(), []int{2}
}

func (x *ListItemsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ListItemsRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *ListItemsRequest) GetPagination() *common.PaginationParam {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items *item.OrderItems `protobuf:"bytes,1,opt,name=items,proto3" json:"items,omitempty"`
}

func (x *ListItemsResponse) Reset() {
	*x = ListItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_item_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsResponse) ProtoMessage() {}

func (x *ListItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_item_item_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsResponse.ProtoReflect.Descriptor instead.
func (*ListItemsResponse) Descriptor() ([]byte, []int) {
	return file_api_order_item_item_proto_rawDescGZIP(), []int{3}
}

func (x *ListItemsResponse) GetItems() *item.OrderItems {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreatItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string          `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Item    *item.OrderItem `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *CreatItemRequest) Reset() {
	*x = CreatItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_item_item_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatItemRequest) ProtoMessage() {}

func (x *CreatItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_item_item_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatItemRequest.ProtoReflect.Descriptor instead.
func (*CreatItemRequest) Descriptor() ([]byte, []int) {
	return file_api_order_item_item_proto_rawDescGZIP(), []int{4}
}

func (x *CreatItemRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreatItemRequest) GetItem() *item.OrderItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type CreatItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreatItemResponse) Reset() {
	*x = CreatItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_item_item_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatItemResponse) ProtoMessage() {}

func (x *CreatItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_item_item_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatItemResponse.ProtoReflect.Descriptor instead.
func (*CreatItemResponse) Descriptor() ([]byte, []int) {
	return file_api_order_item_item_proto_rawDescGZIP(), []int{5}
}

func (x *CreatItemResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId string          `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Item    *item.OrderItem `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *UpdateItemRequest) Reset() {
	*x = UpdateItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_item_item_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemRequest) ProtoMessage() {}

func (x *UpdateItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_item_item_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemRequest.ProtoReflect.Descriptor instead.
func (*UpdateItemRequest) Descriptor() ([]byte, []int) {
	return file_api_order_item_item_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateItemRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *UpdateItemRequest) GetItem() *item.OrderItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type UpdateItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updated int64 `protobuf:"varint,1,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *UpdateItemResponse) Reset() {
	*x = UpdateItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_item_item_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemResponse) ProtoMessage() {}

func (x *UpdateItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_item_item_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemResponse.ProtoReflect.Descriptor instead.
func (*UpdateItemResponse) Descriptor() ([]byte, []int) {
	return file_api_order_item_item_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateItemResponse) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

type DeleteItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId string `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *DeleteItemRequest) Reset() {
	*x = DeleteItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_item_item_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemRequest) ProtoMessage() {}

func (x *DeleteItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_item_item_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemRequest.ProtoReflect.Descriptor instead.
func (*DeleteItemRequest) Descriptor() ([]byte, []int) {
	return file_api_order_item_item_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteItemRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type DeleteItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted int64 `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeleteItemResponse) Reset() {
	*x = DeleteItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_item_item_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemResponse) ProtoMessage() {}

func (x *DeleteItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_item_item_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemResponse.ProtoReflect.Descriptor instead.
func (*DeleteItemResponse) Descriptor() ([]byte, []int) {
	return file_api_order_item_item_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteItemResponse) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

var File_api_order_item_item_proto protoreflect.FileDescriptor

var file_api_order_item_item_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x69, 0x74, 0x65, 0x6d,
	0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x70, 0x6b, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x69, 0x74, 0x65,
	0x6d, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x8c,
	0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x4b, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4b, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x62, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x23,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x73, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x2e, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x3e, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x32, 0xc3, 0x05, 0x0a, 0x0b, 0x49, 0x74, 0x65,
	0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x24, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x12, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x62, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x12, 0x87, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x69, 0x74, 0x65,
	0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x28, 0x12, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x62, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x85, 0x01, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22,
	0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x3a, 0x01, 0x2a, 0x12, 0x8c, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x27,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x1a, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a,
	0x01, 0x2a, 0x12, 0x89, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x27, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x69, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x2a, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x7d, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x19,
	0x5a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f,
	0x69, 0x74, 0x65, 0x6d, 0x3b, 0x69, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_order_item_item_proto_rawDescOnce sync.Once
	file_api_order_item_item_proto_rawDescData = file_api_order_item_item_proto_rawDesc
)

func file_api_order_item_item_proto_rawDescGZIP() []byte {
	file_api_order_item_item_proto_rawDescOnce.Do(func() {
		file_api_order_item_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_order_item_item_proto_rawDescData)
	})
	return file_api_order_item_item_proto_rawDescData
}

var file_api_order_item_item_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_order_item_item_proto_goTypes = []interface{}{
	(*GetItemRequest)(nil),         // 0: pkg.proto.order.item.GetItemRequest
	(*GetItemResponse)(nil),        // 1: pkg.proto.order.item.GetItemResponse
	(*ListItemsRequest)(nil),       // 2: pkg.proto.order.item.ListItemsRequest
	(*ListItemsResponse)(nil),      // 3: pkg.proto.order.item.ListItemsResponse
	(*CreatItemRequest)(nil),       // 4: pkg.proto.order.item.CreatItemRequest
	(*CreatItemResponse)(nil),      // 5: pkg.proto.order.item.CreatItemResponse
	(*UpdateItemRequest)(nil),      // 6: pkg.proto.order.item.UpdateItemRequest
	(*UpdateItemResponse)(nil),     // 7: pkg.proto.order.item.UpdateItemResponse
	(*DeleteItemRequest)(nil),      // 8: pkg.proto.order.item.DeleteItemRequest
	(*DeleteItemResponse)(nil),     // 9: pkg.proto.order.item.DeleteItemResponse
	(*item.OrderItem)(nil),         // 10: pkg.proto.order.item.OrderItem
	(*common.PaginationParam)(nil), // 11: pkg.proto.common.PaginationParam
	(*item.OrderItems)(nil),        // 12: pkg.proto.order.item.OrderItems
}
var file_api_order_item_item_proto_depIdxs = []int32{
	10, // 0: pkg.proto.order.item.GetItemResponse.item:type_name -> pkg.proto.order.item.OrderItem
	11, // 1: pkg.proto.order.item.ListItemsRequest.pagination:type_name -> pkg.proto.common.PaginationParam
	12, // 2: pkg.proto.order.item.ListItemsResponse.items:type_name -> pkg.proto.order.item.OrderItems
	10, // 3: pkg.proto.order.item.CreatItemRequest.item:type_name -> pkg.proto.order.item.OrderItem
	10, // 4: pkg.proto.order.item.UpdateItemRequest.item:type_name -> pkg.proto.order.item.OrderItem
	0,  // 5: pkg.proto.order.item.ItemService.Get:input_type -> pkg.proto.order.item.GetItemRequest
	2,  // 6: pkg.proto.order.item.ItemService.List:input_type -> pkg.proto.order.item.ListItemsRequest
	4,  // 7: pkg.proto.order.item.ItemService.Create:input_type -> pkg.proto.order.item.CreatItemRequest
	6,  // 8: pkg.proto.order.item.ItemService.Update:input_type -> pkg.proto.order.item.UpdateItemRequest
	8,  // 9: pkg.proto.order.item.ItemService.Delete:input_type -> pkg.proto.order.item.DeleteItemRequest
	1,  // 10: pkg.proto.order.item.ItemService.Get:output_type -> pkg.proto.order.item.GetItemResponse
	3,  // 11: pkg.proto.order.item.ItemService.List:output_type -> pkg.proto.order.item.ListItemsResponse
	5,  // 12: pkg.proto.order.item.ItemService.Create:output_type -> pkg.proto.order.item.CreatItemResponse
	7,  // 13: pkg.proto.order.item.ItemService.Update:output_type -> pkg.proto.order.item.UpdateItemResponse
	9,  // 14: pkg.proto.order.item.ItemService.Delete:output_type -> pkg.proto.order.item.DeleteItemResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_order_item_item_proto_init() }
func file_api_order_item_item_proto_init() {
	if File_api_order_item_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_order_item_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_order_item_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_order_item_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_order_item_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_order_item_item_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatItemRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_order_item_item_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatItemResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_order_item_item_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_order_item_item_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_order_item_item_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_order_item_item_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_order_item_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_order_item_item_proto_goTypes,
		DependencyIndexes: file_api_order_item_item_proto_depIdxs,
		MessageInfos:      file_api_order_item_item_proto_msgTypes,
	}.Build()
	File_api_order_item_item_proto = out.File
	file_api_order_item_item_proto_rawDesc = nil
	file_api_order_item_item_proto_goTypes = nil
	file_api_order_item_item_proto_depIdxs = nil
}
