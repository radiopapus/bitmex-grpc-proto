// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: utility.proto

package bitmex_grpc_proto

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
	UtilityService_Ping_FullMethodName        = "/bitmex.UtilityService/Ping"
	UtilityService_SetLogLevel_FullMethodName = "/bitmex.UtilityService/SetLogLevel"
)

// UtilityServiceClient is the client API for UtilityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// utility сервис для служебных нужд
type UtilityServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	SetLogLevel(ctx context.Context, in *LogLevelRequest, opts ...grpc.CallOption) (*LogLevelResponse, error)
}

type utilityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUtilityServiceClient(cc grpc.ClientConnInterface) UtilityServiceClient {
	return &utilityServiceClient{cc}
}

func (c *utilityServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, UtilityService_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *utilityServiceClient) SetLogLevel(ctx context.Context, in *LogLevelRequest, opts ...grpc.CallOption) (*LogLevelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogLevelResponse)
	err := c.cc.Invoke(ctx, UtilityService_SetLogLevel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UtilityServiceServer is the server API for UtilityService service.
// All implementations must embed UnimplementedUtilityServiceServer
// for forward compatibility.
//
// utility сервис для служебных нужд
type UtilityServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	SetLogLevel(context.Context, *LogLevelRequest) (*LogLevelResponse, error)
	mustEmbedUnimplementedUtilityServiceServer()
}

// UnimplementedUtilityServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUtilityServiceServer struct{}

func (UnimplementedUtilityServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedUtilityServiceServer) SetLogLevel(context.Context, *LogLevelRequest) (*LogLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLogLevel not implemented")
}
func (UnimplementedUtilityServiceServer) mustEmbedUnimplementedUtilityServiceServer() {}
func (UnimplementedUtilityServiceServer) testEmbeddedByValue()                        {}

// UnsafeUtilityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UtilityServiceServer will
// result in compilation errors.
type UnsafeUtilityServiceServer interface {
	mustEmbedUnimplementedUtilityServiceServer()
}

func RegisterUtilityServiceServer(s grpc.ServiceRegistrar, srv UtilityServiceServer) {
	// If the following call pancis, it indicates UnimplementedUtilityServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UtilityService_ServiceDesc, srv)
}

func _UtilityService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilityServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UtilityService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilityServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UtilityService_SetLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilityServiceServer).SetLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UtilityService_SetLogLevel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilityServiceServer).SetLogLevel(ctx, req.(*LogLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UtilityService_ServiceDesc is the grpc.ServiceDesc for UtilityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UtilityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bitmex.UtilityService",
	HandlerType: (*UtilityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _UtilityService_Ping_Handler,
		},
		{
			MethodName: "SetLogLevel",
			Handler:    _UtilityService_SetLogLevel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "utility.proto",
}
