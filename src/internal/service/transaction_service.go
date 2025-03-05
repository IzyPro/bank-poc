package service

import (
	"bank-poc/src/internal/domain"
	utils "bank-poc/src/internal/utils/enums"
	"log"
	"time"

	"github.com/google/uuid"
)

type TransactionService struct {
	repo domain.TransactionRepository
}

func NewTransactionService(repo *domain.TransactionRepository) *TransactionService {
	return &TransactionService{
		repo: *repo,
	}
}

func (service *TransactionService) Deposit(transactionDto *domain.TransactionDto) *domain.ApiResponse {
	t := &domain.Transaction{
		Id:              uuid.New().String(),
		CreatedAt:       time.Now(),
		Surcharge:       transactionDto.Surcharge,
		Amount:          transactionDto.Amount,
		Reference:       transactionDto.Reference,
		Narration:       transactionDto.Narration,
		TransactionType: utils.Credit,
	}

	res := new(domain.ApiResponse)

	// Validation
	valid := t.IsValid()
	if valid != nil {
		log.Println(valid)
		return res.Failure(valid.Error())
	}
	return service.repo.Deposit(t)
}

func (service *TransactionService) Withdraw(transactionDto *domain.TransactionDto) *domain.ApiResponse {
	t := &domain.Transaction{
		Id:              uuid.New().String(),
		CreatedAt:       time.Now(),
		Surcharge:       transactionDto.Surcharge,
		Amount:          transactionDto.Amount,
		Reference:       transactionDto.Reference,
		Narration:       transactionDto.Narration,
		TransactionType: utils.Debit,
	}

	res := new(domain.ApiResponse)

	// Validation
	valid := t.IsValid()
	if valid != nil {
		log.Println(valid)
		return res.Failure(valid.Error())
	}
	return service.repo.Withdraw(t)
}

func (service *TransactionService) Balance() *domain.ApiResponse {
	res := new(domain.ApiResponse)

	balance := service.repo.Balance()
	if !balance.Successful || balance.Data == nil {
		return balance
	}

	return res.Success("Account balance retrieved successfully", balance.Data)
}

func (service *TransactionService) TransactionHistory() *domain.ApiResponse {
	return service.repo.TransactionHistory()
}

func (service *TransactionService) Rollback() *domain.ApiResponse {
	return service.repo.Rollback()
}
