// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package category

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryServiceClient interface {
	// Get category
	Get(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*GetCategoryResponse, error)
	// List categories
	List(ctx context.Context, in *ListCategoriesRequest, opts ...grpc.CallOption) (*ListCategoriesResponse, error)
	// Create category
	Create(ctx context.Context, in *CreatCategoryRequest, opts ...grpc.CallOption) (*CreatCategoryResponse, error)
	// Update category
	Update(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*UpdateCategoryResponse, error)
	// Delete category
	Delete(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) Get(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*GetCategoryResponse, error) {
	out := new(GetCategoryResponse)
	err := c.cc.Invoke(ctx, "/pkg.proto.category.CategoryService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) List(ctx context.Context, in *ListCategoriesRequest, opts ...grpc.CallOption) (*ListCategoriesResponse, error) {
	out := new(ListCategoriesResponse)
	err := c.cc.Invoke(ctx, "/pkg.proto.category.CategoryService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) Create(ctx context.Context, in *CreatCategoryRequest, opts ...grpc.CallOption) (*CreatCategoryResponse, error) {
	out := new(CreatCategoryResponse)
	err := c.cc.Invoke(ctx, "/pkg.proto.category.CategoryService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) Update(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*UpdateCategoryResponse, error) {
	out := new(UpdateCategoryResponse)
	err := c.cc.Invoke(ctx, "/pkg.proto.category.CategoryService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) Delete(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error) {
	out := new(DeleteCategoryResponse)
	err := c.cc.Invoke(ctx, "/pkg.proto.category.CategoryService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServiceServer is the server API for CategoryService service.
// All implementations should embed UnimplementedCategoryServiceServer
// for forward compatibility
type CategoryServiceServer interface {
	// Get category
	Get(context.Context, *GetCategoryRequest) (*GetCategoryResponse, error)
	// List categories
	List(context.Context, *ListCategoriesRequest) (*ListCategoriesResponse, error)
	// Create category
	Create(context.Context, *CreatCategoryRequest) (*CreatCategoryResponse, error)
	// Update category
	Update(context.Context, *UpdateCategoryRequest) (*UpdateCategoryResponse, error)
	// Delete category
	Delete(context.Context, *DeleteCategoryRequest) (*DeleteCategoryResponse, error)
}

// UnimplementedCategoryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCategoryServiceServer struct {
}

func (UnimplementedCategoryServiceServer) Get(context.Context, *GetCategoryRequest) (*GetCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCategoryServiceServer) List(context.Context, *ListCategoriesRequest) (*ListCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCategoryServiceServer) Create(context.Context, *CreatCategoryRequest) (*CreatCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCategoryServiceServer) Update(context.Context, *UpdateCategoryRequest) (*UpdateCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCategoryServiceServer) Delete(context.Context, *DeleteCategoryRequest) (*DeleteCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServiceServer will
// result in compilation errors.
type UnsafeCategoryServiceServer interface {
	mustEmbedUnimplementedCategoryServiceServer()
}

func RegisterCategoryServiceServer(s *grpc.Server, srv CategoryServiceServer) {
	s.RegisterService(&_CategoryService_serviceDesc, srv)
}

func _CategoryService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.proto.category.CategoryService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Get(ctx, req.(*GetCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.proto.category.CategoryService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).List(ctx, req.(*ListCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.proto.category.CategoryService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Create(ctx, req.(*CreatCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.proto.category.CategoryService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Update(ctx, req.(*UpdateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.proto.category.CategoryService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Delete(ctx, req.(*DeleteCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CategoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pkg.proto.category.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CategoryService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CategoryService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CategoryService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CategoryService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CategoryService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/category/category.proto",
}
