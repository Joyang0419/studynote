package transaction

import (
	"context"
	"fmt"
)

func NewTransactionServer() TransactionServer {
	return new(server)
}

type UserID = string
type Balance = float64

var (
	UserBalanceMap = map[UserID]Balance{
		"1": 10000,
	}
)

type server struct {
	UnimplementedTransactionServer // 必須得embedded
}

func (s *server) Withdraw(ctx context.Context, request *WithdrawRequest) (*TransactionFormatResponse, error) {
	balance, exists := UserBalanceMap[request.UserID]
	if !exists {
		return nil, fmt.Errorf("server Balance UserBalanceMap !exists, userID: %s", request.UserID)
	}

	calculate := balance - request.Amount
	if calculate < 0 {
		return nil, fmt.Errorf("server Withdraw balance is not enough, balance: %v, amount: %v", balance, request.Amount)
	}

	UserBalanceMap[request.UserID] = calculate

	return &TransactionFormatResponse{Balance: calculate}, nil
}

func (s *server) Deposits(ctx context.Context, request *DepositsRequest) (*TransactionFormatResponse, error) {
	balance, exists := UserBalanceMap[request.UserID]
	if !exists {
		return nil, fmt.Errorf("server Balance UserBalanceMap !exists, userID: %s", request.UserID)
	}

	calculate := balance + request.Amount
	UserBalanceMap[request.UserID] = calculate
	return &TransactionFormatResponse{Balance: calculate}, nil
}

func (s *server) Balance(ctx context.Context, user *User) (*TransactionFormatResponse, error) {
	balance, exists := UserBalanceMap[user.UserID]
	if !exists {
		return nil, fmt.Errorf("server Balance UserBalanceMap !exists, userID: %s", user.UserID)
	}

	return &TransactionFormatResponse{
		Balance: balance,
	}, nil
}
