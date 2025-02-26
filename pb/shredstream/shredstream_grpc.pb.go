// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: shredstream.proto

package shredstream

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
	Shredstream_SendHeartbeat_FullMethodName = "/shredstream.Shredstream/SendHeartbeat"
)

// ShredstreamClient is the client API for Shredstream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShredstreamClient interface {
	// RPC endpoint to send heartbeats to keep shreds flowing
	SendHeartbeat(ctx context.Context, in *Heartbeat, opts ...grpc.CallOption) (*HeartbeatResponse, error)
}

type shredstreamClient struct {
	cc grpc.ClientConnInterface
}

func NewShredstreamClient(cc grpc.ClientConnInterface) ShredstreamClient {
	return &shredstreamClient{cc}
}

func (c *shredstreamClient) SendHeartbeat(ctx context.Context, in *Heartbeat, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, Shredstream_SendHeartbeat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShredstreamServer is the server API for Shredstream service.
// All implementations must embed UnimplementedShredstreamServer
// for forward compatibility.
type ShredstreamServer interface {
	// RPC endpoint to send heartbeats to keep shreds flowing
	SendHeartbeat(context.Context, *Heartbeat) (*HeartbeatResponse, error)
	mustEmbedUnimplementedShredstreamServer()
}

// UnimplementedShredstreamServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedShredstreamServer struct{}

func (UnimplementedShredstreamServer) SendHeartbeat(context.Context, *Heartbeat) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendHeartbeat not implemented")
}
func (UnimplementedShredstreamServer) mustEmbedUnimplementedShredstreamServer() {}
func (UnimplementedShredstreamServer) testEmbeddedByValue()                     {}

// UnsafeShredstreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShredstreamServer will
// result in compilation errors.
type UnsafeShredstreamServer interface {
	mustEmbedUnimplementedShredstreamServer()
}

func RegisterShredstreamServer(s grpc.ServiceRegistrar, srv ShredstreamServer) {
	// If the following call pancis, it indicates UnimplementedShredstreamServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Shredstream_ServiceDesc, srv)
}

func _Shredstream_SendHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Heartbeat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShredstreamServer).SendHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shredstream_SendHeartbeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShredstreamServer).SendHeartbeat(ctx, req.(*Heartbeat))
	}
	return interceptor(ctx, in, info, handler)
}

// Shredstream_ServiceDesc is the grpc.ServiceDesc for Shredstream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shredstream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shredstream.Shredstream",
	HandlerType: (*ShredstreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendHeartbeat",
			Handler:    _Shredstream_SendHeartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shredstream.proto",
}
