// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CircleClient is the client API for Circle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CircleClient interface {
	Area(ctx context.Context, in *AreaRequest, opts ...grpc.CallOption) (*AreaResponse, error)
}

type circleClient struct {
	cc grpc.ClientConnInterface
}

func NewCircleClient(cc grpc.ClientConnInterface) CircleClient {
	return &circleClient{cc}
}

func (c *circleClient) Area(ctx context.Context, in *AreaRequest, opts ...grpc.CallOption) (*AreaResponse, error) {
	out := new(AreaResponse)
	err := c.cc.Invoke(ctx, "/api.Circle/Area", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CircleServer is the server API for Circle service.
// All implementations must embed UnimplementedCircleServer
// for forward compatibility
type CircleServer interface {
	Area(context.Context, *AreaRequest) (*AreaResponse, error)
	mustEmbedUnimplementedCircleServer()
}

// UnimplementedCircleServer must be embedded to have forward compatible implementations.
type UnimplementedCircleServer struct {
}

func (UnimplementedCircleServer) Area(context.Context, *AreaRequest) (*AreaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Area not implemented")
}
func (UnimplementedCircleServer) mustEmbedUnimplementedCircleServer() {}

// UnsafeCircleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CircleServer will
// result in compilation errors.
type UnsafeCircleServer interface {
	mustEmbedUnimplementedCircleServer()
}

func RegisterCircleServer(s *grpc.Server, srv CircleServer) {
	s.RegisterService(&_Circle_serviceDesc, srv)
}

func _Circle_Area_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CircleServer).Area(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Circle/Area",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CircleServer).Area(ctx, req.(*AreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Circle_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Circle",
	HandlerType: (*CircleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Area",
			Handler:    _Circle_Area_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}

// MathClient is the client API for Math service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MathClient interface {
	Sqr(ctx context.Context, in *SqrRequest, opts ...grpc.CallOption) (*SqrResponse, error)
}

type mathClient struct {
	cc grpc.ClientConnInterface
}

func NewMathClient(cc grpc.ClientConnInterface) MathClient {
	return &mathClient{cc}
}

func (c *mathClient) Sqr(ctx context.Context, in *SqrRequest, opts ...grpc.CallOption) (*SqrResponse, error) {
	out := new(SqrResponse)
	err := c.cc.Invoke(ctx, "/api.Math/Sqr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathServer is the server API for Math service.
// All implementations must embed UnimplementedMathServer
// for forward compatibility
type MathServer interface {
	Sqr(context.Context, *SqrRequest) (*SqrResponse, error)
	mustEmbedUnimplementedMathServer()
}

// UnimplementedMathServer must be embedded to have forward compatible implementations.
type UnimplementedMathServer struct {
}

func (UnimplementedMathServer) Sqr(context.Context, *SqrRequest) (*SqrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sqr not implemented")
}
func (UnimplementedMathServer) mustEmbedUnimplementedMathServer() {}

// UnsafeMathServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MathServer will
// result in compilation errors.
type UnsafeMathServer interface {
	mustEmbedUnimplementedMathServer()
}

func RegisterMathServer(s *grpc.Server, srv MathServer) {
	s.RegisterService(&_Math_serviceDesc, srv)
}

func _Math_Sqr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServer).Sqr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Math/Sqr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServer).Sqr(ctx, req.(*SqrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Math_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Math",
	HandlerType: (*MathServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sqr",
			Handler:    _Math_Sqr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}
