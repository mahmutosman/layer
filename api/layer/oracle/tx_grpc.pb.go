// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package oracle

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
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	SubmitValue(ctx context.Context, in *MsgSubmitValue, opts ...grpc.CallOption) (*MsgSubmitValueResponse, error)
	CommitReport(ctx context.Context, in *MsgCommitReport, opts ...grpc.CallOption) (*MsgCommitReportResponse, error)
	Tip(ctx context.Context, in *MsgTip, opts ...grpc.CallOption) (*MsgTipResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/layer.oracle.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitValue(ctx context.Context, in *MsgSubmitValue, opts ...grpc.CallOption) (*MsgSubmitValueResponse, error) {
	out := new(MsgSubmitValueResponse)
	err := c.cc.Invoke(ctx, "/layer.oracle.Msg/SubmitValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CommitReport(ctx context.Context, in *MsgCommitReport, opts ...grpc.CallOption) (*MsgCommitReportResponse, error) {
	out := new(MsgCommitReportResponse)
	err := c.cc.Invoke(ctx, "/layer.oracle.Msg/CommitReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Tip(ctx context.Context, in *MsgTip, opts ...grpc.CallOption) (*MsgTipResponse, error) {
	out := new(MsgTipResponse)
	err := c.cc.Invoke(ctx, "/layer.oracle.Msg/Tip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	SubmitValue(context.Context, *MsgSubmitValue) (*MsgSubmitValueResponse, error)
	CommitReport(context.Context, *MsgCommitReport) (*MsgCommitReportResponse, error)
	Tip(context.Context, *MsgTip) (*MsgTipResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) SubmitValue(context.Context, *MsgSubmitValue) (*MsgSubmitValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitValue not implemented")
}
func (UnimplementedMsgServer) CommitReport(context.Context, *MsgCommitReport) (*MsgCommitReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitReport not implemented")
}
func (UnimplementedMsgServer) Tip(context.Context, *MsgTip) (*MsgTipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tip not implemented")
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

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/layer.oracle.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/layer.oracle.Msg/SubmitValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitValue(ctx, req.(*MsgSubmitValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CommitReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCommitReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CommitReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/layer.oracle.Msg/CommitReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CommitReport(ctx, req.(*MsgCommitReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Tip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTip)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Tip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/layer.oracle.Msg/Tip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Tip(ctx, req.(*MsgTip))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "layer.oracle.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "SubmitValue",
			Handler:    _Msg_SubmitValue_Handler,
		},
		{
			MethodName: "CommitReport",
			Handler:    _Msg_CommitReport_Handler,
		},
		{
			MethodName: "Tip",
			Handler:    _Msg_Tip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "layer/oracle/tx.proto",
}
