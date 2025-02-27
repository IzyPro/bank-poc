package repository

import (
	"bank-poc/src/internal/domain"
	"sync"
)

type InMemoryTransactionRepository struct {
	// DB connection and other fields
}

func NewTransactionRepository() *InMemoryTransactionRepository {
	return &InMemoryTransactionRepository{}
}

var account domain.Account = domain.Account{Balance: 0, Transactions: nil}
var mu sync.Mutex

func (repo *InMemoryTransactionRepository) Deposit(t domain.Transaction) domain.ApiResponse {
	res := new(domain.ApiResponse)

	mu.Lock()
	account.Balance += (t.Amount - t.Surcharge)
	mu.Unlock()
	account.Transactions = append(account.Transactions, t)

	return res.Success("Deposit successful", t)
}

func (repo *InMemoryTransactionRepository) Withdraw(t domain.Transaction) domain.ApiResponse {
	res := new(domain.ApiResponse)

	debitAmount := t.Amount + t.Surcharge
	if account.Balance < debitAmount {
		return res.Failure("Insufficient funds")
	}
	mu.Lock()
	account.Balance -= debitAmount
	mu.Unlock()
	account.Transactions = append(account.Transactions, t)

	return res.Success("Withdrawal successful", t)
}

func (repo *InMemoryTransactionRepository) Balance() domain.ApiResponse {
	res := new(domain.ApiResponse)
	return res.Success("Account Balance retrieved successfully", account.Balance)
}

func (repo *InMemoryTransactionRepository) TransactionHistory() domain.ApiResponse {
	res := new(domain.ApiResponse)

	return res.Success("Transaction history retrieved successfully", account.Transactions)
}
