// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: auction.proto

package auction

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

// ServerNodeClient is the client API for ServerNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerNodeClient interface {
	RequestAccess(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Empty, error)
	RequestLamportTimestamp(ctx context.Context, in *Request, opts ...grpc.CallOption) (*LamportTimestamp, error)
	RequestPeerList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*PeerList, error)
	LetPeerKnowIExist(ctx context.Context, in *ServerNodeInfo, opts ...grpc.CallOption) (*Empty, error)
	Bid(ctx context.Context, in *BidMessage, opts ...grpc.CallOption) (*Ack, error)
	Result(ctx context.Context, in *ResultMessage, opts ...grpc.CallOption) (*Outcome, error)
}

type serverNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewServerNodeClient(cc grpc.ClientConnInterface) ServerNodeClient {
	return &serverNodeClient{cc}
}

func (c *serverNodeClient) RequestAccess(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chat.serverNode/RequestAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverNodeClient) RequestLamportTimestamp(ctx context.Context, in *Request, opts ...grpc.CallOption) (*LamportTimestamp, error) {
	out := new(LamportTimestamp)
	err := c.cc.Invoke(ctx, "/chat.serverNode/RequestLamportTimestamp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverNodeClient) RequestPeerList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*PeerList, error) {
	out := new(PeerList)
	err := c.cc.Invoke(ctx, "/chat.serverNode/RequestPeerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverNodeClient) LetPeerKnowIExist(ctx context.Context, in *ServerNodeInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chat.serverNode/LetPeerKnowIExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverNodeClient) Bid(ctx context.Context, in *BidMessage, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/chat.serverNode/bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverNodeClient) Result(ctx context.Context, in *ResultMessage, opts ...grpc.CallOption) (*Outcome, error) {
	out := new(Outcome)
	err := c.cc.Invoke(ctx, "/chat.serverNode/result", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerNodeServer is the server API for ServerNode service.
// All implementations must embed UnimplementedServerNodeServer
// for forward compatibility
type ServerNodeServer interface {
	RequestAccess(context.Context, *Request) (*Empty, error)
	RequestLamportTimestamp(context.Context, *Request) (*LamportTimestamp, error)
	RequestPeerList(context.Context, *Request) (*PeerList, error)
	LetPeerKnowIExist(context.Context, *ServerNodeInfo) (*Empty, error)
	Bid(context.Context, *BidMessage) (*Ack, error)
	Result(context.Context, *ResultMessage) (*Outcome, error)
	mustEmbedUnimplementedServerNodeServer()
}

// UnimplementedServerNodeServer must be embedded to have forward compatible implementations.
type UnimplementedServerNodeServer struct {
}

func (UnimplementedServerNodeServer) RequestAccess(context.Context, *Request) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAccess not implemented")
}
func (UnimplementedServerNodeServer) RequestLamportTimestamp(context.Context, *Request) (*LamportTimestamp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestLamportTimestamp not implemented")
}
func (UnimplementedServerNodeServer) RequestPeerList(context.Context, *Request) (*PeerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestPeerList not implemented")
}
func (UnimplementedServerNodeServer) LetPeerKnowIExist(context.Context, *ServerNodeInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LetPeerKnowIExist not implemented")
}
func (UnimplementedServerNodeServer) Bid(context.Context, *BidMessage) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedServerNodeServer) Result(context.Context, *ResultMessage) (*Outcome, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedServerNodeServer) mustEmbedUnimplementedServerNodeServer() {}

// UnsafeServerNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerNodeServer will
// result in compilation errors.
type UnsafeServerNodeServer interface {
	mustEmbedUnimplementedServerNodeServer()
}

func RegisterServerNodeServer(s grpc.ServiceRegistrar, srv ServerNodeServer) {
	s.RegisterService(&ServerNode_ServiceDesc, srv)
}

func _ServerNode_RequestAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerNodeServer).RequestAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.serverNode/RequestAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerNodeServer).RequestAccess(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerNode_RequestLamportTimestamp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerNodeServer).RequestLamportTimestamp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.serverNode/RequestLamportTimestamp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerNodeServer).RequestLamportTimestamp(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerNode_RequestPeerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerNodeServer).RequestPeerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.serverNode/RequestPeerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerNodeServer).RequestPeerList(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerNode_LetPeerKnowIExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerNodeInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerNodeServer).LetPeerKnowIExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.serverNode/LetPeerKnowIExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerNodeServer).LetPeerKnowIExist(ctx, req.(*ServerNodeInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerNode_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerNodeServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.serverNode/bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerNodeServer).Bid(ctx, req.(*BidMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerNode_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerNodeServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.serverNode/result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerNodeServer).Result(ctx, req.(*ResultMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerNode_ServiceDesc is the grpc.ServiceDesc for ServerNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.serverNode",
	HandlerType: (*ServerNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestAccess",
			Handler:    _ServerNode_RequestAccess_Handler,
		},
		{
			MethodName: "RequestLamportTimestamp",
			Handler:    _ServerNode_RequestLamportTimestamp_Handler,
		},
		{
			MethodName: "RequestPeerList",
			Handler:    _ServerNode_RequestPeerList_Handler,
		},
		{
			MethodName: "LetPeerKnowIExist",
			Handler:    _ServerNode_LetPeerKnowIExist_Handler,
		},
		{
			MethodName: "bid",
			Handler:    _ServerNode_Bid_Handler,
		},
		{
			MethodName: "result",
			Handler:    _ServerNode_Result_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auction.proto",
}
