// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: protopkg/service.proto

package protopkg

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PriceFetcher_FetchPrice_FullMethodName = "/protopkg.PriceFetcher/FetchPrice"
)

// PriceFetcherClient is the client API for PriceFetcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PriceFetcherClient interface {
	FetchPrice(ctx context.Context, in *FetchPriceRequest, opts ...grpc.CallOption) (*FetchPriceResponse, error)
}

type priceFetcherClient struct {
	cc grpc.ClientConnInterface
}

func NewPriceFetcherClient(cc grpc.ClientConnInterface) PriceFetcherClient {
	return &priceFetcherClient{cc}
}

func (c *priceFetcherClient) FetchPrice(ctx context.Context, in *FetchPriceRequest, opts ...grpc.CallOption) (*FetchPriceResponse, error) {
	out := new(FetchPriceResponse)
	err := c.cc.Invoke(ctx, PriceFetcher_FetchPrice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PriceFetcherServer is the server API for PriceFetcher service.
// All implementations must embed UnimplementedPriceFetcherServer
// for forward compatibility
type PriceFetcherServer interface {
	FetchPrice(context.Context, *FetchPriceRequest) (*FetchPriceResponse, error)
	mustEmbedUnimplementedPriceFetcherServer()
}

// UnimplementedPriceFetcherServer must be embedded to have forward compatible implementations.
type UnimplementedPriceFetcherServer struct {
}

func (UnimplementedPriceFetcherServer) FetchPrice(context.Context, *FetchPriceRequest) (*FetchPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchPrice not implemented")
}
func (UnimplementedPriceFetcherServer) mustEmbedUnimplementedPriceFetcherServer() {}

// UnsafePriceFetcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PriceFetcherServer will
// result in compilation errors.
type UnsafePriceFetcherServer interface {
	mustEmbedUnimplementedPriceFetcherServer()
}

func RegisterPriceFetcherServer(s grpc.ServiceRegistrar, srv PriceFetcherServer) {
	s.RegisterService(&PriceFetcher_ServiceDesc, srv)
}

func _PriceFetcher_FetchPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceFetcherServer).FetchPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PriceFetcher_FetchPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceFetcherServer).FetchPrice(ctx, req.(*FetchPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PriceFetcher_ServiceDesc is the grpc.ServiceDesc for PriceFetcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PriceFetcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protopkg.PriceFetcher",
	HandlerType: (*PriceFetcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchPrice",
			Handler:    _PriceFetcher_FetchPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protopkg/service.proto",
}
