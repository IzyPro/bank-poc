package repository

import "bank-poc/src/internal/domain"

type Repository struct {
	TransactionRepo domain.TransactionRepository
}

func New() *Repository {
	// Return all Repository Implementations
	return &Repository{
		TransactionRepo: NewTransactionRepository(),
	}
}
