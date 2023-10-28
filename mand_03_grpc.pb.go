// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mand_03.proto

package helloworld

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

// ChittychatClient is the client API for Chittychat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChittychatClient interface {
	// Sends a greeting
	ConnectToServer(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionReply, error)
	Broadcast(ctx context.Context, in *BroadcastMessage, opts ...grpc.CallOption) (*Confirmation, error)
	Publish(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Confirmation, error)
	Disconnect(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Confirmation, error)
}

type chittychatClient struct {
	cc grpc.ClientConnInterface
}

func NewChittychatClient(cc grpc.ClientConnInterface) ChittychatClient {
	return &chittychatClient{cc}
}

func (c *chittychatClient) ConnectToServer(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionReply, error) {
	out := new(ConnectionReply)
	err := c.cc.Invoke(ctx, "/mand_03.Chittychat/ConnectToServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittychatClient) Broadcast(ctx context.Context, in *BroadcastMessage, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := c.cc.Invoke(ctx, "/mand_03.Chittychat/broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittychatClient) Publish(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := c.cc.Invoke(ctx, "/mand_03.Chittychat/publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittychatClient) Disconnect(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := c.cc.Invoke(ctx, "/mand_03.Chittychat/disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChittychatServer is the server API for Chittychat service.
// All implementations must embed UnimplementedChittychatServer
// for forward compatibility
type ChittychatServer interface {
	// Sends a greeting
	ConnectToServer(context.Context, *ConnectionRequest) (*ConnectionReply, error)
	Broadcast(context.Context, *BroadcastMessage) (*Confirmation, error)
	Publish(context.Context, *Msg) (*Confirmation, error)
	Disconnect(context.Context, *Msg) (*Confirmation, error)
	mustEmbedUnimplementedChittychatServer()
}

// UnimplementedChittychatServer must be embedded to have forward compatible implementations.
type UnimplementedChittychatServer struct {
}

func (UnimplementedChittychatServer) ConnectToServer(context.Context, *ConnectionRequest) (*ConnectionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectToServer not implemented")
}
func (UnimplementedChittychatServer) Broadcast(context.Context, *BroadcastMessage) (*Confirmation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedChittychatServer) Publish(context.Context, *Msg) (*Confirmation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedChittychatServer) Disconnect(context.Context, *Msg) (*Confirmation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedChittychatServer) mustEmbedUnimplementedChittychatServer() {}

// UnsafeChittychatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChittychatServer will
// result in compilation errors.
type UnsafeChittychatServer interface {
	mustEmbedUnimplementedChittychatServer()
}

func RegisterChittychatServer(s grpc.ServiceRegistrar, srv ChittychatServer) {
	s.RegisterService(&Chittychat_ServiceDesc, srv)
}

func _Chittychat_ConnectToServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittychatServer).ConnectToServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mand_03.Chittychat/ConnectToServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittychatServer).ConnectToServer(ctx, req.(*ConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chittychat_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittychatServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mand_03.Chittychat/broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittychatServer).Broadcast(ctx, req.(*BroadcastMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chittychat_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittychatServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mand_03.Chittychat/publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittychatServer).Publish(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chittychat_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittychatServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mand_03.Chittychat/disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittychatServer).Disconnect(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

// Chittychat_ServiceDesc is the grpc.ServiceDesc for Chittychat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chittychat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mand_03.Chittychat",
	HandlerType: (*ChittychatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConnectToServer",
			Handler:    _Chittychat_ConnectToServer_Handler,
		},
		{
			MethodName: "broadcast",
			Handler:    _Chittychat_Broadcast_Handler,
		},
		{
			MethodName: "publish",
			Handler:    _Chittychat_Publish_Handler,
		},
		{
			MethodName: "disconnect",
			Handler:    _Chittychat_Disconnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mand_03.proto",
}
