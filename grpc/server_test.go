package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"studynote/grpc/service/transaction"
)

/*
TestNewGrpcServer, 這是啟動Grpc server,
下面這三個是Grpc Client的應用(必須得先啟動server):
- TestNewGrpcServerBalance
- TestNewGrpcServerDeposits
- TestNewGrpcServerWithdraw
*/

func TestNewGrpcServer(t *testing.T) {
	grpcServer := NewGrpcServer(ServerRouter{Transaction: transaction.NewTransactionServer()})
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func TestNewGrpcServerBalance(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	transactionClient := transaction.NewTransactionClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := transactionClient.Balance(ctx, &transaction.User{UserID: "1"})
	if err != nil {
		log.Fatalf("transactionClient.Balance err: %v", err)
	}
	log.Printf("Balance: %v", response.Balance)
}

func TestNewGrpcServerDeposits(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	transactionClient := transaction.NewTransactionClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := transactionClient.Deposits(ctx, &transaction.DepositsRequest{
		UserID: "1",
		Amount: 10000,
	},
	)
	if err != nil {
		log.Fatalf("transactionClient.Deposits err: %v", err)
	}
	log.Printf("Balance: %v", response.Balance)
}

func TestNewGrpcServerWithdraw(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	transactionClient := transaction.NewTransactionClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := transactionClient.Withdraw(ctx, &transaction.WithdrawRequest{
		UserID: "1",
		Amount: 10000,
	},
	)
	if err != nil {
		log.Fatalf("transactionClient.Withdraw err: %v", err)
	}
	log.Printf("Balance: %v", response.Balance)
}
