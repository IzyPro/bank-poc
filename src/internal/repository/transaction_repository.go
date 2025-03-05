package repository

import (
	"bank-poc/src/internal/domain"
	utils "bank-poc/src/internal/utils/enums"
	"sync"
)

type InMemoryTransactionRepository struct {
	// DB connection and other fields
}

func NewTransactionRepository() *InMemoryTransactionRepository {
	return &InMemoryTransactionRepository{}
}

var account domain.Account = domain.Account{Balance: 0, Transactions: nil}
var mutex sync.Mutex

func (repo *InMemoryTransactionRepository) Deposit(trnx *domain.Transaction) *domain.ApiResponse {
	res := new(domain.ApiResponse)

	mutex.Lock()
	account.Balance += (trnx.Amount - trnx.Surcharge)
	mutex.Unlock()
	account.Transactions = append(account.Transactions, *trnx)

	return res.Success("Deposit successful", trnx)
}

func (repo *InMemoryTransactionRepository) Withdraw(trnx *domain.Transaction) *domain.ApiResponse {
	res := new(domain.ApiResponse)

	debitAmount := trnx.Amount + trnx.Surcharge

	mutex.Lock()
	// Retrieve balance from db
	if account.Balance < debitAmount {
		return res.Failure("Insufficient funds")
	}

	account.Balance -= debitAmount
	mutex.Unlock()
	account.Transactions = append(account.Transactions, *trnx)

	return res.Success("Withdrawal successful", trnx)
}

func (repo *InMemoryTransactionRepository) Balance() *domain.ApiResponse {
	res := new(domain.ApiResponse)
	return res.Success("Account Balance retrieved successfully", account.Balance)
}

func (repo *InMemoryTransactionRepository) TransactionHistory() *domain.ApiResponse {
	res := new(domain.ApiResponse)
	return res.Success("Transaction history retrieved successfully", account.Transactions)
}

func (repo *InMemoryTransactionRepository) Rollback() *domain.ApiResponse {
	res := new(domain.ApiResponse)

	if len(account.Transactions) < 1 {
		return res.Failure("Transaction list is empty")
	}

	lastTrnxIndex := len(account.Transactions) - 1
	trnx := account.Transactions[lastTrnxIndex]

	account.Transactions = account.Transactions[:lastTrnxIndex]
	if trnx.TransactionType == utils.Credit {
		account.Balance -= (trnx.Amount - trnx.Surcharge)
	} else {
		account.Balance += (trnx.Amount + trnx.Surcharge)
	}

	return res.Success("Rollback successful", trnx)
}
