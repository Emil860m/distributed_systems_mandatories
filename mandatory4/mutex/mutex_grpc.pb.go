// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: mutex.proto

package mutex

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

// MutexClient is the client API for Mutex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MutexClient interface {
	RequestAccess(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Empty, error)
	RequestLamportTimestamp(ctx context.Context, in *Request, opts ...grpc.CallOption) (*LamportTimestamp, error)
	RequestPeerList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*PeerList, error)
	LetPeerKnowIExist(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*Empty, error)
}

type mutexClient struct {
	cc grpc.ClientConnInterface
}

func NewMutexClient(cc grpc.ClientConnInterface) MutexClient {
	return &mutexClient{cc}
}

func (c *mutexClient) RequestAccess(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chat.Mutex/RequestAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mutexClient) RequestLamportTimestamp(ctx context.Context, in *Request, opts ...grpc.CallOption) (*LamportTimestamp, error) {
	out := new(LamportTimestamp)
	err := c.cc.Invoke(ctx, "/chat.Mutex/RequestLamportTimestamp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mutexClient) RequestPeerList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*PeerList, error) {
	out := new(PeerList)
	err := c.cc.Invoke(ctx, "/chat.Mutex/RequestPeerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mutexClient) LetPeerKnowIExist(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chat.Mutex/LetPeerKnowIExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MutexServer is the server API for Mutex service.
// All implementations must embed UnimplementedMutexServer
// for forward compatibility
type MutexServer interface {
	RequestAccess(context.Context, *Request) (*Empty, error)
	RequestLamportTimestamp(context.Context, *Request) (*LamportTimestamp, error)
	RequestPeerList(context.Context, *Request) (*PeerList, error)
	LetPeerKnowIExist(context.Context, *ClientInfo) (*Empty, error)
	mustEmbedUnimplementedMutexServer()
}

// UnimplementedMutexServer must be embedded to have forward compatible implementations.
type UnimplementedMutexServer struct {
}

func (UnimplementedMutexServer) RequestAccess(context.Context, *Request) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAccess not implemented")
}
func (UnimplementedMutexServer) RequestLamportTimestamp(context.Context, *Request) (*LamportTimestamp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestLamportTimestamp not implemented")
}
func (UnimplementedMutexServer) RequestPeerList(context.Context, *Request) (*PeerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestPeerList not implemented")
}
func (UnimplementedMutexServer) LetPeerKnowIExist(context.Context, *ClientInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LetPeerKnowIExist not implemented")
}
func (UnimplementedMutexServer) mustEmbedUnimplementedMutexServer() {}

// UnsafeMutexServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MutexServer will
// result in compilation errors.
type UnsafeMutexServer interface {
	mustEmbedUnimplementedMutexServer()
}

func RegisterMutexServer(s grpc.ServiceRegistrar, srv MutexServer) {
	s.RegisterService(&Mutex_ServiceDesc, srv)
}

func _Mutex_RequestAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutexServer).RequestAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Mutex/RequestAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutexServer).RequestAccess(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mutex_RequestLamportTimestamp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutexServer).RequestLamportTimestamp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Mutex/RequestLamportTimestamp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutexServer).RequestLamportTimestamp(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mutex_RequestPeerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutexServer).RequestPeerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Mutex/RequestPeerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutexServer).RequestPeerList(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mutex_LetPeerKnowIExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutexServer).LetPeerKnowIExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Mutex/LetPeerKnowIExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutexServer).LetPeerKnowIExist(ctx, req.(*ClientInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// Mutex_ServiceDesc is the grpc.ServiceDesc for Mutex service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mutex_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Mutex",
	HandlerType: (*MutexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestAccess",
			Handler:    _Mutex_RequestAccess_Handler,
		},
		{
			MethodName: "RequestLamportTimestamp",
			Handler:    _Mutex_RequestLamportTimestamp_Handler,
		},
		{
			MethodName: "RequestPeerList",
			Handler:    _Mutex_RequestPeerList_Handler,
		},
		{
			MethodName: "LetPeerKnowIExist",
			Handler:    _Mutex_LetPeerKnowIExist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mutex.proto",
}
