// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PimServiceClient is the client API for PimService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PimServiceClient interface {
	SearchProducts(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*Products, error)
	GetAllCategoriesByShop(ctx context.Context, in *ShopID, opts ...grpc.CallOption) (*CategoriesTrees, error)
	CreateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetProduct(ctx context.Context, in *IDs, opts ...grpc.CallOption) (*Product, error)
	UpdateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteProduct(ctx context.Context, in *IDs, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCategory(ctx context.Context, in *IDs, opts ...grpc.CallOption) (*Category, error)
	UpdateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteCategory(ctx context.Context, in *IDs, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type pimServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPimServiceClient(cc grpc.ClientConnInterface) PimServiceClient {
	return &pimServiceClient{cc}
}

func (c *pimServiceClient) SearchProducts(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*Products, error) {
	out := new(Products)
	err := c.cc.Invoke(ctx, "/pim_service.PimService/SearchProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pimServiceClient) GetAllCategoriesByShop(ctx context.Context, in *ShopID, opts ...grpc.CallOption) (*CategoriesTrees, error) {
	out := new(CategoriesTrees)
	err := c.cc.Invoke(ctx, "/pim_service.PimService/GetAllCategoriesByShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pimServiceClient) CreateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pim_service.PimService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pimServiceClient) GetProduct(ctx context.Context, in *IDs, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/pim_service.PimService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pimServiceClient) UpdateProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pim_service.PimService/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pimServiceClient) DeleteProduct(ctx context.Context, in *IDs, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pim_service.PimService/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pimServiceClient) CreateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pim_service.PimService/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pimServiceClient) GetCategory(ctx context.Context, in *IDs, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/pim_service.PimService/GetCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pimServiceClient) UpdateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pim_service.PimService/UpdateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pimServiceClient) DeleteCategory(ctx context.Context, in *IDs, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pim_service.PimService/DeleteCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PimServiceServer is the server API for PimService service.
// All implementations must embed UnimplementedPimServiceServer
// for forward compatibility
type PimServiceServer interface {
	SearchProducts(context.Context, *SearchRequest) (*Products, error)
	GetAllCategoriesByShop(context.Context, *ShopID) (*CategoriesTrees, error)
	CreateProduct(context.Context, *Product) (*emptypb.Empty, error)
	GetProduct(context.Context, *IDs) (*Product, error)
	UpdateProduct(context.Context, *Product) (*emptypb.Empty, error)
	DeleteProduct(context.Context, *IDs) (*emptypb.Empty, error)
	CreateCategory(context.Context, *Category) (*emptypb.Empty, error)
	GetCategory(context.Context, *IDs) (*Category, error)
	UpdateCategory(context.Context, *Category) (*emptypb.Empty, error)
	DeleteCategory(context.Context, *IDs) (*emptypb.Empty, error)
	mustEmbedUnimplementedPimServiceServer()
}

// UnimplementedPimServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPimServiceServer struct {
}

func (UnimplementedPimServiceServer) SearchProducts(context.Context, *SearchRequest) (*Products, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProducts not implemented")
}
func (UnimplementedPimServiceServer) GetAllCategoriesByShop(context.Context, *ShopID) (*CategoriesTrees, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategoriesByShop not implemented")
}
func (UnimplementedPimServiceServer) CreateProduct(context.Context, *Product) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedPimServiceServer) GetProduct(context.Context, *IDs) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedPimServiceServer) UpdateProduct(context.Context, *Product) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedPimServiceServer) DeleteProduct(context.Context, *IDs) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedPimServiceServer) CreateCategory(context.Context, *Category) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedPimServiceServer) GetCategory(context.Context, *IDs) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedPimServiceServer) UpdateCategory(context.Context, *Category) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedPimServiceServer) DeleteCategory(context.Context, *IDs) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedPimServiceServer) mustEmbedUnimplementedPimServiceServer() {}

// UnsafePimServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PimServiceServer will
// result in compilation errors.
type UnsafePimServiceServer interface {
	mustEmbedUnimplementedPimServiceServer()
}

func RegisterPimServiceServer(s grpc.ServiceRegistrar, srv PimServiceServer) {
	s.RegisterService(&PimService_ServiceDesc, srv)
}

func _PimService_SearchProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PimServiceServer).SearchProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pim_service.PimService/SearchProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PimServiceServer).SearchProducts(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PimService_GetAllCategoriesByShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PimServiceServer).GetAllCategoriesByShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pim_service.PimService/GetAllCategoriesByShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PimServiceServer).GetAllCategoriesByShop(ctx, req.(*ShopID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PimService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PimServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pim_service.PimService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PimServiceServer).CreateProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _PimService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PimServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pim_service.PimService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PimServiceServer).GetProduct(ctx, req.(*IDs))
	}
	return interceptor(ctx, in, info, handler)
}

func _PimService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PimServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pim_service.PimService/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PimServiceServer).UpdateProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _PimService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PimServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pim_service.PimService/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PimServiceServer).DeleteProduct(ctx, req.(*IDs))
	}
	return interceptor(ctx, in, info, handler)
}

func _PimService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PimServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pim_service.PimService/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PimServiceServer).CreateCategory(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _PimService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PimServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pim_service.PimService/GetCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PimServiceServer).GetCategory(ctx, req.(*IDs))
	}
	return interceptor(ctx, in, info, handler)
}

func _PimService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PimServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pim_service.PimService/UpdateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PimServiceServer).UpdateCategory(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _PimService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PimServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pim_service.PimService/DeleteCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PimServiceServer).DeleteCategory(ctx, req.(*IDs))
	}
	return interceptor(ctx, in, info, handler)
}

// PimService_ServiceDesc is the grpc.ServiceDesc for PimService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PimService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pim_service.PimService",
	HandlerType: (*PimServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchProducts",
			Handler:    _PimService_SearchProducts_Handler,
		},
		{
			MethodName: "GetAllCategoriesByShop",
			Handler:    _PimService_GetAllCategoriesByShop_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _PimService_CreateProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _PimService_GetProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _PimService_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _PimService_DeleteProduct_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _PimService_CreateCategory_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _PimService_GetCategory_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _PimService_UpdateCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _PimService_DeleteCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pim-service/pim-service.proto",
}