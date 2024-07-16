// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: warden/act/v1beta1/tx.proto

package actv1beta1

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

const (
	Msg_UpdateParams_FullMethodName  = "/warden.act.v1beta1.Msg/UpdateParams"
	Msg_NewAction_FullMethodName     = "/warden.act.v1beta1.Msg/NewAction"
	Msg_ApproveAction_FullMethodName = "/warden.act.v1beta1.Msg/ApproveAction"
	Msg_CheckAction_FullMethodName   = "/warden.act.v1beta1.Msg/CheckAction"
	Msg_NewRule_FullMethodName       = "/warden.act.v1beta1.Msg/NewRule"
	Msg_UpdateRule_FullMethodName    = "/warden.act.v1beta1.Msg/UpdateRule"
	Msg_RevokeAction_FullMethodName  = "/warden.act.v1beta1.Msg/RevokeAction"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	// NewAction creates a new Action.
	NewAction(ctx context.Context, in *MsgNewAction, opts ...grpc.CallOption) (*MsgNewActionResponse, error)
	// Add an approval to an existing Action.
	ApproveAction(ctx context.Context, in *MsgApproveAction, opts ...grpc.CallOption) (*MsgApproveActionResponse, error)
	// Add an approval to an existing Action.
	CheckAction(ctx context.Context, in *MsgCheckAction, opts ...grpc.CallOption) (*MsgCheckActionResponse, error)
	// Create a new Rule.
	NewRule(ctx context.Context, in *MsgNewRule, opts ...grpc.CallOption) (*MsgNewRuleResponse, error)
	// Update an existing act name and definition.
	UpdateRule(ctx context.Context, in *MsgUpdateRule, opts ...grpc.CallOption) (*MsgUpdateRuleResponse, error)
	// Revoke an existing Action while in pending state.
	RevokeAction(ctx context.Context, in *MsgRevokeAction, opts ...grpc.CallOption) (*MsgRevokeActionResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) NewAction(ctx context.Context, in *MsgNewAction, opts ...grpc.CallOption) (*MsgNewActionResponse, error) {
	out := new(MsgNewActionResponse)
	err := c.cc.Invoke(ctx, Msg_NewAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ApproveAction(ctx context.Context, in *MsgApproveAction, opts ...grpc.CallOption) (*MsgApproveActionResponse, error) {
	out := new(MsgApproveActionResponse)
	err := c.cc.Invoke(ctx, Msg_ApproveAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CheckAction(ctx context.Context, in *MsgCheckAction, opts ...grpc.CallOption) (*MsgCheckActionResponse, error) {
	out := new(MsgCheckActionResponse)
	err := c.cc.Invoke(ctx, Msg_CheckAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) NewRule(ctx context.Context, in *MsgNewRule, opts ...grpc.CallOption) (*MsgNewRuleResponse, error) {
	out := new(MsgNewRuleResponse)
	err := c.cc.Invoke(ctx, Msg_NewRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateRule(ctx context.Context, in *MsgUpdateRule, opts ...grpc.CallOption) (*MsgUpdateRuleResponse, error) {
	out := new(MsgUpdateRuleResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RevokeAction(ctx context.Context, in *MsgRevokeAction, opts ...grpc.CallOption) (*MsgRevokeActionResponse, error) {
	out := new(MsgRevokeActionResponse)
	err := c.cc.Invoke(ctx, Msg_RevokeAction_FullMethodName, in, out, opts...)
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
	// NewAction creates a new Action.
	NewAction(context.Context, *MsgNewAction) (*MsgNewActionResponse, error)
	// Add an approval to an existing Action.
	ApproveAction(context.Context, *MsgApproveAction) (*MsgApproveActionResponse, error)
	// Add an approval to an existing Action.
	CheckAction(context.Context, *MsgCheckAction) (*MsgCheckActionResponse, error)
	// Create a new Rule.
	NewRule(context.Context, *MsgNewRule) (*MsgNewRuleResponse, error)
	// Update an existing act name and definition.
	UpdateRule(context.Context, *MsgUpdateRule) (*MsgUpdateRuleResponse, error)
	// Revoke an existing Action while in pending state.
	RevokeAction(context.Context, *MsgRevokeAction) (*MsgRevokeActionResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) NewAction(context.Context, *MsgNewAction) (*MsgNewActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAction not implemented")
}
func (UnimplementedMsgServer) ApproveAction(context.Context, *MsgApproveAction) (*MsgApproveActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveAction not implemented")
}
func (UnimplementedMsgServer) CheckAction(context.Context, *MsgCheckAction) (*MsgCheckActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAction not implemented")
}
func (UnimplementedMsgServer) NewRule(context.Context, *MsgNewRule) (*MsgNewRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewRule not implemented")
}
func (UnimplementedMsgServer) UpdateRule(context.Context, *MsgUpdateRule) (*MsgUpdateRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRule not implemented")
}
func (UnimplementedMsgServer) RevokeAction(context.Context, *MsgRevokeAction) (*MsgRevokeActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeAction not implemented")
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
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_NewAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgNewAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).NewAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_NewAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).NewAction(ctx, req.(*MsgNewAction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ApproveAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgApproveAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ApproveAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ApproveAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ApproveAction(ctx, req.(*MsgApproveAction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CheckAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCheckAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CheckAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CheckAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CheckAction(ctx, req.(*MsgCheckAction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_NewRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgNewRule)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).NewRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_NewRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).NewRule(ctx, req.(*MsgNewRule))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateRule)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateRule(ctx, req.(*MsgUpdateRule))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RevokeAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRevokeAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RevokeAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RevokeAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RevokeAction(ctx, req.(*MsgRevokeAction))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "warden.act.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "NewAction",
			Handler:    _Msg_NewAction_Handler,
		},
		{
			MethodName: "ApproveAction",
			Handler:    _Msg_ApproveAction_Handler,
		},
		{
			MethodName: "CheckAction",
			Handler:    _Msg_CheckAction_Handler,
		},
		{
			MethodName: "NewRule",
			Handler:    _Msg_NewRule_Handler,
		},
		{
			MethodName: "UpdateRule",
			Handler:    _Msg_UpdateRule_Handler,
		},
		{
			MethodName: "RevokeAction",
			Handler:    _Msg_RevokeAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warden/act/v1beta1/tx.proto",
}
