/// Please use the following editor setup for this file:

// Tab size=2; Tabs as spaces; Clean up trailing whitepsace
// 'make proto' will run clang-format to fix formatiing

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: v1/watchdog/watchdogService.proto

package watchdog

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
	WatchDog_Reload_FullMethodName = "/watchdog.WatchDog/Reload"
)

// WatchDogClient is the client API for WatchDog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WatchDogClient interface {
	// when input domain list has modified
	// reload the new data from yaml file
	// and rerun the domain checks
	Reload(ctx context.Context, in *ReloadRequest, opts ...grpc.CallOption) (*ReloadResponse, error)
}

type watchDogClient struct {
	cc grpc.ClientConnInterface
}

func NewWatchDogClient(cc grpc.ClientConnInterface) WatchDogClient {
	return &watchDogClient{cc}
}

func (c *watchDogClient) Reload(ctx context.Context, in *ReloadRequest, opts ...grpc.CallOption) (*ReloadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReloadResponse)
	err := c.cc.Invoke(ctx, WatchDog_Reload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WatchDogServer is the server API for WatchDog service.
// All implementations must embed UnimplementedWatchDogServer
// for forward compatibility.
type WatchDogServer interface {
	// when input domain list has modified
	// reload the new data from yaml file
	// and rerun the domain checks
	Reload(context.Context, *ReloadRequest) (*ReloadResponse, error)
	mustEmbedUnimplementedWatchDogServer()
}

// UnimplementedWatchDogServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWatchDogServer struct{}

func (UnimplementedWatchDogServer) Reload(context.Context, *ReloadRequest) (*ReloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reload not implemented")
}
func (UnimplementedWatchDogServer) mustEmbedUnimplementedWatchDogServer() {}
func (UnimplementedWatchDogServer) testEmbeddedByValue()                  {}

// UnsafeWatchDogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WatchDogServer will
// result in compilation errors.
type UnsafeWatchDogServer interface {
	mustEmbedUnimplementedWatchDogServer()
}

func RegisterWatchDogServer(s grpc.ServiceRegistrar, srv WatchDogServer) {
	// If the following call pancis, it indicates UnimplementedWatchDogServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WatchDog_ServiceDesc, srv)
}

func _WatchDog_Reload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatchDogServer).Reload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WatchDog_Reload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatchDogServer).Reload(ctx, req.(*ReloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WatchDog_ServiceDesc is the grpc.ServiceDesc for WatchDog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WatchDog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "watchdog.WatchDog",
	HandlerType: (*WatchDogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reload",
			Handler:    _WatchDog_Reload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/watchdog/watchdogService.proto",
}
