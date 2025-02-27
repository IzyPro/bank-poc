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

func (s *TransactionService) Deposit(transactionDto domain.TransactionDto) domain.ApiResponse {
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
	return s.repo.Deposit(*t)
}

func (s *TransactionService) Withdraw(transactionDto domain.TransactionDto) domain.ApiResponse {
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
	return s.repo.Withdraw(*t)
}

func (s *TransactionService) Balance() domain.ApiResponse {
	res := new(domain.ApiResponse)

	balance := s.repo.Balance()
	if !balance.Successful || balance.Data == nil {
		return balance
	}

	// account_balance.balance = fmt.Sprintf("%f", balance.Data)
	// account_balance.timestamp = time.Now()

	return res.Success("Account balance retrieved successfully", balance.Data)
}

func (s *TransactionService) TransactionHistory() domain.ApiResponse {
	return s.repo.TransactionHistory()
}
