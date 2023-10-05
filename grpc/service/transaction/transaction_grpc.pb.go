// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.1
// source: grpc/service/transaction/transaction.proto

package transaction

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
	Transaction_Withdraw_FullMethodName = "/transaction.Transaction/Withdraw"
	Transaction_Deposits_FullMethodName = "/transaction.Transaction/Deposits"
	Transaction_Balance_FullMethodName  = "/transaction.Transaction/Balance"
)

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionClient interface {
	Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*TransactionFormatResponse, error)
	Deposits(ctx context.Context, in *DepositsRequest, opts ...grpc.CallOption) (*TransactionFormatResponse, error)
	Balance(ctx context.Context, in *User, opts ...grpc.CallOption) (*TransactionFormatResponse, error)
}

type transactionClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionClient(cc grpc.ClientConnInterface) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*TransactionFormatResponse, error) {
	out := new(TransactionFormatResponse)
	err := c.cc.Invoke(ctx, Transaction_Withdraw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) Deposits(ctx context.Context, in *DepositsRequest, opts ...grpc.CallOption) (*TransactionFormatResponse, error) {
	out := new(TransactionFormatResponse)
	err := c.cc.Invoke(ctx, Transaction_Deposits_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) Balance(ctx context.Context, in *User, opts ...grpc.CallOption) (*TransactionFormatResponse, error) {
	out := new(TransactionFormatResponse)
	err := c.cc.Invoke(ctx, Transaction_Balance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
// All implementations must embed UnimplementedTransactionServer
// for forward compatibility
type TransactionServer interface {
	Withdraw(context.Context, *WithdrawRequest) (*TransactionFormatResponse, error)
	Deposits(context.Context, *DepositsRequest) (*TransactionFormatResponse, error)
	Balance(context.Context, *User) (*TransactionFormatResponse, error)
	mustEmbedUnimplementedTransactionServer()
}

// UnimplementedTransactionServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (UnimplementedTransactionServer) Withdraw(context.Context, *WithdrawRequest) (*TransactionFormatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedTransactionServer) Deposits(context.Context, *DepositsRequest) (*TransactionFormatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposits not implemented")
}
func (UnimplementedTransactionServer) Balance(context.Context, *User) (*TransactionFormatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (UnimplementedTransactionServer) mustEmbedUnimplementedTransactionServer() {}

// UnsafeTransactionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServer will
// result in compilation errors.
type UnsafeTransactionServer interface {
	mustEmbedUnimplementedTransactionServer()
}

func RegisterTransactionServer(s grpc.ServiceRegistrar, srv TransactionServer) {
	s.RegisterService(&Transaction_ServiceDesc, srv)
}

func _Transaction_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transaction_Withdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).Withdraw(ctx, req.(*WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_Deposits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).Deposits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transaction_Deposits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).Deposits(ctx, req.(*DepositsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transaction_Balance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).Balance(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// Transaction_ServiceDesc is the grpc.ServiceDesc for Transaction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transaction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Withdraw",
			Handler:    _Transaction_Withdraw_Handler,
		},
		{
			MethodName: "Deposits",
			Handler:    _Transaction_Deposits_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _Transaction_Balance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/service/transaction/transaction.proto",
}
