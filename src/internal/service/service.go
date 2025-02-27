package service

import (
	"bank-poc/src/internal/domain"
	"bank-poc/src/internal/repository"
)

type Service struct {
	TransactionService domain.TransactionService
}

func New(repo *repository.Repository) *Service {
	// Return all Service Implementations
	return &Service{
		TransactionService: NewTransactionService(&repo.TransactionRepo),
	}
}
