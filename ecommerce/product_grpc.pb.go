// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: ecommerce/product.proto

package ecommerce

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProductInfoClient is the client API for ProductInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductInfoClient interface {
	AddProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductID, error)
	GetProduct(ctx context.Context, in *ProductID, opts ...grpc.CallOption) (*Product, error)
	GreetUser(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*wrappers.StringValue, error)
}

type productInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewProductInfoClient(cc grpc.ClientConnInterface) ProductInfoClient {
	return &productInfoClient{cc}
}

func (c *productInfoClient) AddProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductID, error) {
	out := new(ProductID)
	err := c.cc.Invoke(ctx, "/productservice.ProductInfo/addProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productInfoClient) GetProduct(ctx context.Context, in *ProductID, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/productservice.ProductInfo/getProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productInfoClient) GreetUser(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*wrappers.StringValue, error) {
	out := new(wrappers.StringValue)
	err := c.cc.Invoke(ctx, "/productservice.ProductInfo/greetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductInfoServer is the server API for ProductInfo service.
// All implementations must embed UnimplementedProductInfoServer
// for forward compatibility
type ProductInfoServer interface {
	AddProduct(context.Context, *Product) (*ProductID, error)
	GetProduct(context.Context, *ProductID) (*Product, error)
	GreetUser(context.Context, *empty.Empty) (*wrappers.StringValue, error)
	mustEmbedUnimplementedProductInfoServer()
}

// UnimplementedProductInfoServer must be embedded to have forward compatible implementations.
type UnimplementedProductInfoServer struct {
}

func (UnimplementedProductInfoServer) AddProduct(context.Context, *Product) (*ProductID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (UnimplementedProductInfoServer) GetProduct(context.Context, *ProductID) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductInfoServer) GreetUser(context.Context, *empty.Empty) (*wrappers.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetUser not implemented")
}
func (UnimplementedProductInfoServer) mustEmbedUnimplementedProductInfoServer() {}

// UnsafeProductInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductInfoServer will
// result in compilation errors.
type UnsafeProductInfoServer interface {
	mustEmbedUnimplementedProductInfoServer()
}

func RegisterProductInfoServer(s grpc.ServiceRegistrar, srv ProductInfoServer) {
	s.RegisterService(&ProductInfo_ServiceDesc, srv)
}

func _ProductInfo_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductInfoServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productservice.ProductInfo/addProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductInfoServer).AddProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductInfo_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductInfoServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productservice.ProductInfo/getProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductInfoServer).GetProduct(ctx, req.(*ProductID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductInfo_GreetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductInfoServer).GreetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productservice.ProductInfo/greetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductInfoServer).GreetUser(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductInfo_ServiceDesc is the grpc.ServiceDesc for ProductInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "productservice.ProductInfo",
	HandlerType: (*ProductInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addProduct",
			Handler:    _ProductInfo_AddProduct_Handler,
		},
		{
			MethodName: "getProduct",
			Handler:    _ProductInfo_GetProduct_Handler,
		},
		{
			MethodName: "greetUser",
			Handler:    _ProductInfo_GreetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ecommerce/product.proto",
}
