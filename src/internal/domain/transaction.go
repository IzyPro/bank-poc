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

func (trnx Transaction) IsValid() error {
	if len(trnx.Reference) < 1 {
		return errors.New("reference cannot be empty")
	}
	if trnx.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if trnx.Surcharge < 0 {
		return errors.New("surcharge must be greater than zero")
	}
	if trnx.Amount <= trnx.Surcharge {
		return errors.New("amount must be greater than transaction fee")
	}

	return nil
}

type TransactionRepository interface {
	Deposit(trnx *Transaction) *ApiResponse
	Withdraw(trnx *Transaction) *ApiResponse
	Balance() *ApiResponse
	TransactionHistory() *ApiResponse
	Rollback() *ApiResponse
}

type TransactionService interface {
	Deposit(trnx *TransactionDto) *ApiResponse
	Withdraw(trnx *TransactionDto) *ApiResponse
	Balance() *ApiResponse
	TransactionHistory() *ApiResponse
	Rollback() *ApiResponse
}
