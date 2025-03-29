package repository

import "github.com/fbpr/mnc-test-go/internal/domain"

type CustomerRepository interface {
	GetByEmail(email string) (*domain.Customer, error)
	UpdateLoginStatus(email string, isLoggedIn bool) error
}
