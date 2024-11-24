// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: tickers.proto

package grpc_gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TickersService_GetMultipleTickers_FullMethodName = "/TickersService/GetMultipleTickers"
)

// TickersServiceClient is the client API for TickersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TickersServiceClient interface {
	GetMultipleTickers(ctx context.Context, in *TickersRequest, opts ...grpc.CallOption) (*MultipleTickerResponse, error)
}

type tickersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTickersServiceClient(cc grpc.ClientConnInterface) TickersServiceClient {
	return &tickersServiceClient{cc}
}

func (c *tickersServiceClient) GetMultipleTickers(ctx context.Context, in *TickersRequest, opts ...grpc.CallOption) (*MultipleTickerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MultipleTickerResponse)
	err := c.cc.Invoke(ctx, TickersService_GetMultipleTickers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TickersServiceServer is the server API for TickersService service.
// All implementations must embed UnimplementedTickersServiceServer
// for forward compatibility.
type TickersServiceServer interface {
	GetMultipleTickers(context.Context, *TickersRequest) (*MultipleTickerResponse, error)
	mustEmbedUnimplementedTickersServiceServer()
}

// UnimplementedTickersServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTickersServiceServer struct{}

func (UnimplementedTickersServiceServer) GetMultipleTickers(context.Context, *TickersRequest) (*MultipleTickerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultipleTickers not implemented")
}
func (UnimplementedTickersServiceServer) mustEmbedUnimplementedTickersServiceServer() {}
func (UnimplementedTickersServiceServer) testEmbeddedByValue()                        {}

// UnsafeTickersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TickersServiceServer will
// result in compilation errors.
type UnsafeTickersServiceServer interface {
	mustEmbedUnimplementedTickersServiceServer()
}

func RegisterTickersServiceServer(s grpc.ServiceRegistrar, srv TickersServiceServer) {
	// If the following call pancis, it indicates UnimplementedTickersServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TickersService_ServiceDesc, srv)
}

func _TickersService_GetMultipleTickers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TickersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TickersServiceServer).GetMultipleTickers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TickersService_GetMultipleTickers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TickersServiceServer).GetMultipleTickers(ctx, req.(*TickersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TickersService_ServiceDesc is the grpc.ServiceDesc for TickersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TickersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TickersService",
	HandlerType: (*TickersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMultipleTickers",
			Handler:    _TickersService_GetMultipleTickers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tickers.proto",
}