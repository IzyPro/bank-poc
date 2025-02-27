package domain

import (
	utils "bank-poc/src/internal/utils/enums"
	"errors"
	"time"
)

type TransactionDto struct {
	Reference string
	Amount    float64
	Surcharge float64
	Narration string
}

type Transaction struct {
	Id              string
	Reference       string
	CreatedAt       time.Time
	Amount          float64
	Surcharge       float64
	Narration       string
	TransactionType utils.TransactionType
}

func (e Transaction) IsValid() error {
	if len(e.Reference) < 1 {
		return errors.New("reference cannot be empty")
	}
	if e.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if e.Amount <= e.Surcharge {
		return errors.New("amount must be greater than transaction fee")
	}

	return nil
}

type TransactionRepository interface {
	Deposit(t Transaction) ApiResponse
	Withdraw(t Transaction) ApiResponse
	Balance() ApiResponse
	TransactionHistory() ApiResponse
}

type TransactionService interface {
	Deposit(t TransactionDto) ApiResponse
	Withdraw(t TransactionDto) ApiResponse
	Balance() ApiResponse
	TransactionHistory() ApiResponse
}
