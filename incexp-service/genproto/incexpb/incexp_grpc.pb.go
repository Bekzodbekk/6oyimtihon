// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: proto/incexp.proto

package incexpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	IncExpService_RegisterIncome_FullMethodName         = "/IncExpService/RegisterIncome"
	IncExpService_RegisterExpense_FullMethodName        = "/IncExpService/RegisterExpense"
	IncExpService_GetListIncomeVSExpense_FullMethodName = "/IncExpService/GetListIncomeVSExpense"
)

// IncExpServiceClient is the client API for IncExpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IncExpServiceClient interface {
	RegisterIncome(ctx context.Context, in *RegisterIncomeRequest, opts ...grpc.CallOption) (*RegisterIncomeResponse, error)
	RegisterExpense(ctx context.Context, in *RegisterExpenseRequest, opts ...grpc.CallOption) (*RegisterExpenseResponse, error)
	GetListIncomeVSExpense(ctx context.Context, in *GetListIncomeVSExpenseRequest, opts ...grpc.CallOption) (*GetListIncomeVSExpenseResponse, error)
}

type incExpServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIncExpServiceClient(cc grpc.ClientConnInterface) IncExpServiceClient {
	return &incExpServiceClient{cc}
}

func (c *incExpServiceClient) RegisterIncome(ctx context.Context, in *RegisterIncomeRequest, opts ...grpc.CallOption) (*RegisterIncomeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterIncomeResponse)
	err := c.cc.Invoke(ctx, IncExpService_RegisterIncome_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incExpServiceClient) RegisterExpense(ctx context.Context, in *RegisterExpenseRequest, opts ...grpc.CallOption) (*RegisterExpenseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterExpenseResponse)
	err := c.cc.Invoke(ctx, IncExpService_RegisterExpense_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incExpServiceClient) GetListIncomeVSExpense(ctx context.Context, in *GetListIncomeVSExpenseRequest, opts ...grpc.CallOption) (*GetListIncomeVSExpenseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetListIncomeVSExpenseResponse)
	err := c.cc.Invoke(ctx, IncExpService_GetListIncomeVSExpense_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IncExpServiceServer is the server API for IncExpService service.
// All implementations must embed UnimplementedIncExpServiceServer
// for forward compatibility
type IncExpServiceServer interface {
	RegisterIncome(context.Context, *RegisterIncomeRequest) (*RegisterIncomeResponse, error)
	RegisterExpense(context.Context, *RegisterExpenseRequest) (*RegisterExpenseResponse, error)
	GetListIncomeVSExpense(context.Context, *GetListIncomeVSExpenseRequest) (*GetListIncomeVSExpenseResponse, error)
	mustEmbedUnimplementedIncExpServiceServer()
}

// UnimplementedIncExpServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIncExpServiceServer struct {
}

func (UnimplementedIncExpServiceServer) RegisterIncome(context.Context, *RegisterIncomeRequest) (*RegisterIncomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterIncome not implemented")
}
func (UnimplementedIncExpServiceServer) RegisterExpense(context.Context, *RegisterExpenseRequest) (*RegisterExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterExpense not implemented")
}
func (UnimplementedIncExpServiceServer) GetListIncomeVSExpense(context.Context, *GetListIncomeVSExpenseRequest) (*GetListIncomeVSExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListIncomeVSExpense not implemented")
}
func (UnimplementedIncExpServiceServer) mustEmbedUnimplementedIncExpServiceServer() {}

// UnsafeIncExpServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IncExpServiceServer will
// result in compilation errors.
type UnsafeIncExpServiceServer interface {
	mustEmbedUnimplementedIncExpServiceServer()
}

func RegisterIncExpServiceServer(s grpc.ServiceRegistrar, srv IncExpServiceServer) {
	s.RegisterService(&IncExpService_ServiceDesc, srv)
}

func _IncExpService_RegisterIncome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterIncomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncExpServiceServer).RegisterIncome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncExpService_RegisterIncome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncExpServiceServer).RegisterIncome(ctx, req.(*RegisterIncomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncExpService_RegisterExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncExpServiceServer).RegisterExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncExpService_RegisterExpense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncExpServiceServer).RegisterExpense(ctx, req.(*RegisterExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncExpService_GetListIncomeVSExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListIncomeVSExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncExpServiceServer).GetListIncomeVSExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncExpService_GetListIncomeVSExpense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncExpServiceServer).GetListIncomeVSExpense(ctx, req.(*GetListIncomeVSExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IncExpService_ServiceDesc is the grpc.ServiceDesc for IncExpService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IncExpService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "IncExpService",
	HandlerType: (*IncExpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterIncome",
			Handler:    _IncExpService_RegisterIncome_Handler,
		},
		{
			MethodName: "RegisterExpense",
			Handler:    _IncExpService_RegisterExpense_Handler,
		},
		{
			MethodName: "GetListIncomeVSExpense",
			Handler:    _IncExpService_GetListIncomeVSExpense_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/incexp.proto",
}
