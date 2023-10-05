package grpc

import (
	"google.golang.org/grpc"
	"studynote/grpc/service/transaction"
)

type ServerRouter struct {
	Transaction transaction.TransactionServer
}

func NewGrpcServer(router ServerRouter, opt ...grpc.ServerOption) *grpc.Server {
	grpcServer := grpc.NewServer(opt...)
	transaction.RegisterTransactionServer(
		grpcServer,
		router.Transaction,
	)

	return grpcServer
}
