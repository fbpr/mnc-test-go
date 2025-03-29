package repository

import "github.com/fbpr/mnc-test-go/internal/domain"

type TransactionRepository interface {
	GetByID(id string) (domain.Transaction, error)
	UpdateStatus(id, status string) error
}
