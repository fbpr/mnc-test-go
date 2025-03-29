package usecase

import (
	"errors"
	"fmt"

	"github.com/fbpr/mnc-test-go/internal/domain"
	"github.com/fbpr/mnc-test-go/internal/repository"
)

type AuthUseCase struct {
	customerRepository repository.CustomerRepository
	historyRepository  repository.HistoryRepository
}

func NewAuthUseCase(customerRepository repository.CustomerRepository, historyRepository repository.HistoryRepository) *AuthUseCase {
	return &AuthUseCase{
		customerRepository: customerRepository,
		historyRepository:  historyRepository,
	}
}

func (uc *AuthUseCase) Login(email, password string) (domain.LoginResponse, error) {
	customer, err := uc.customerRepository.GetByEmail(email)
	fmt.Println(err)
	if err != nil {
		return domain.LoginResponse{}, errors.New("credentials invalid")
	}

	if customer.Password != password {
		return domain.LoginResponse{}, errors.New("credentials invalid")
	}

	err = uc.customerRepository.UpdateLoginStatus(email, true)
	if err != nil {
		return domain.LoginResponse{}, errors.New("failed to update login status")
	}

	err = uc.historyRepository.CreateLoginHistory(customer.ID)
	if err != nil {
		return domain.LoginResponse{}, fmt.Errorf("failed to create login history")
	}

	response := domain.LoginResponse{
		ID:    customer.ID,
		Email: customer.Email,
		Name:  customer.Name,
	}

	return response, nil
}

func (uc *AuthUseCase) Logout(email string) error {
	customer, err := uc.customerRepository.GetByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	if !customer.Status.IsLoggedIn {
		return errors.New("user not logged in")
	}

	err = uc.customerRepository.UpdateLoginStatus(email, false)
	if err != nil {
		return errors.New("failed to update login status")
	}

	err = uc.historyRepository.CreateLogoutHistory(customer.ID)
	if err != nil {
		return fmt.Errorf("failed to create logout history: %w", err)
	}

	return nil
}
