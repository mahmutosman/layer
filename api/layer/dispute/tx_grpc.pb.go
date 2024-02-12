// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dispute

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	ProposeDispute(ctx context.Context, in *MsgProposeDispute, opts ...grpc.CallOption) (*MsgProposeDisputeResponse, error)
	AddFeeToDispute(ctx context.Context, in *MsgAddFeeToDispute, opts ...grpc.CallOption) (*MsgAddFeeToDisputeResponse, error)
	Vote(ctx context.Context, in *MsgVote, opts ...grpc.CallOption) (*MsgVoteResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ProposeDispute(ctx context.Context, in *MsgProposeDispute, opts ...grpc.CallOption) (*MsgProposeDisputeResponse, error) {
	out := new(MsgProposeDisputeResponse)
	err := c.cc.Invoke(ctx, "/layer.dispute.Msg/ProposeDispute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddFeeToDispute(ctx context.Context, in *MsgAddFeeToDispute, opts ...grpc.CallOption) (*MsgAddFeeToDisputeResponse, error) {
	out := new(MsgAddFeeToDisputeResponse)
	err := c.cc.Invoke(ctx, "/layer.dispute.Msg/AddFeeToDispute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Vote(ctx context.Context, in *MsgVote, opts ...grpc.CallOption) (*MsgVoteResponse, error) {
	out := new(MsgVoteResponse)
	err := c.cc.Invoke(ctx, "/layer.dispute.Msg/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	ProposeDispute(context.Context, *MsgProposeDispute) (*MsgProposeDisputeResponse, error)
	AddFeeToDispute(context.Context, *MsgAddFeeToDispute) (*MsgAddFeeToDisputeResponse, error)
	Vote(context.Context, *MsgVote) (*MsgVoteResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) ProposeDispute(context.Context, *MsgProposeDispute) (*MsgProposeDisputeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProposeDispute not implemented")
}
func (UnimplementedMsgServer) AddFeeToDispute(context.Context, *MsgAddFeeToDispute) (*MsgAddFeeToDisputeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFeeToDispute not implemented")
}
func (UnimplementedMsgServer) Vote(context.Context, *MsgVote) (*MsgVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_ProposeDispute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgProposeDispute)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ProposeDispute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/layer.dispute.Msg/ProposeDispute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ProposeDispute(ctx, req.(*MsgProposeDispute))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddFeeToDispute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddFeeToDispute)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddFeeToDispute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/layer.dispute.Msg/AddFeeToDispute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddFeeToDispute(ctx, req.(*MsgAddFeeToDispute))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/layer.dispute.Msg/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Vote(ctx, req.(*MsgVote))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "layer.dispute.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProposeDispute",
			Handler:    _Msg_ProposeDispute_Handler,
		},
		{
			MethodName: "AddFeeToDispute",
			Handler:    _Msg_AddFeeToDispute_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _Msg_Vote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "layer/dispute/tx.proto",
}
