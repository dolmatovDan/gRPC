// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: currency.proto

package currency

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
	Currency_GetRate_FullMethodName        = "/Currency/GetRate"
	Currency_SubscribeRates_FullMethodName = "/Currency/SubscribeRates"
)

// CurrencyClient is the client API for Currency service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrencyClient interface {
	GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error)
	SubscribeRates(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[RateRequest, StreamingRateResponse], error)
}

type currencyClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyClient(cc grpc.ClientConnInterface) CurrencyClient {
	return &currencyClient{cc}
}

func (c *currencyClient) GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, Currency_GetRate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyClient) SubscribeRates(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[RateRequest, StreamingRateResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Currency_ServiceDesc.Streams[0], Currency_SubscribeRates_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[RateRequest, StreamingRateResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Currency_SubscribeRatesClient = grpc.BidiStreamingClient[RateRequest, StreamingRateResponse]

// CurrencyServer is the server API for Currency service.
// All implementations must embed UnimplementedCurrencyServer
// for forward compatibility.
type CurrencyServer interface {
	GetRate(context.Context, *RateRequest) (*RateResponse, error)
	SubscribeRates(grpc.BidiStreamingServer[RateRequest, StreamingRateResponse]) error
	mustEmbedUnimplementedCurrencyServer()
}

// UnimplementedCurrencyServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCurrencyServer struct{}

func (UnimplementedCurrencyServer) GetRate(context.Context, *RateRequest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRate not implemented")
}
func (UnimplementedCurrencyServer) SubscribeRates(grpc.BidiStreamingServer[RateRequest, StreamingRateResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeRates not implemented")
}
func (UnimplementedCurrencyServer) mustEmbedUnimplementedCurrencyServer() {}
func (UnimplementedCurrencyServer) testEmbeddedByValue()                  {}

// UnsafeCurrencyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrencyServer will
// result in compilation errors.
type UnsafeCurrencyServer interface {
	mustEmbedUnimplementedCurrencyServer()
}

func RegisterCurrencyServer(s grpc.ServiceRegistrar, srv CurrencyServer) {
	// If the following call pancis, it indicates UnimplementedCurrencyServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Currency_ServiceDesc, srv)
}

func _Currency_GetRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServer).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Currency_GetRate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServer).GetRate(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Currency_SubscribeRates_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CurrencyServer).SubscribeRates(&grpc.GenericServerStream[RateRequest, StreamingRateResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Currency_SubscribeRatesServer = grpc.BidiStreamingServer[RateRequest, StreamingRateResponse]

// Currency_ServiceDesc is the grpc.ServiceDesc for Currency service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Currency_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Currency",
	HandlerType: (*CurrencyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRate",
			Handler:    _Currency_GetRate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeRates",
			Handler:       _Currency_SubscribeRates_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "currency.proto",
}
