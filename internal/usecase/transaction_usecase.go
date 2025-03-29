package usecase

import (
	"errors"
	"fmt"

	"github.com/fbpr/mnc-test-go/internal/domain"
	"github.com/fbpr/mnc-test-go/internal/repository"
)

type TransactionUseCase struct {
	transactionRepository repository.TransactionRepository
	customerRepository    repository.CustomerRepository
	historyRepository     repository.HistoryRepository
}

func NewTransactionUseCase(
	transactionRepository repository.TransactionRepository,
	customerRepository repository.CustomerRepository,
	historyRepository repository.HistoryRepository,
) *TransactionUseCase {
	return &TransactionUseCase{
		transactionRepository: transactionRepository,
		customerRepository:    customerRepository,
		historyRepository:     historyRepository,
	}
}

func (uc *TransactionUseCase) ProcessPayment(transactionID string, request domain.PaymentRequest) (domain.PaymentResponse, error) {
	transaction, err := uc.transactionRepository.GetByID(transactionID)
	if err != nil {
		return domain.PaymentResponse{}, err
	}

	customer, err := uc.customerRepository.GetByEmail(request.CustomerEmail)
	if err != nil {
		return domain.PaymentResponse{}, errors.New("customer not found")
	}

	if !customer.Status.IsLoggedIn {
		return domain.PaymentResponse{}, errors.New("customer must be logged in to do the payment")
	}

	if transaction.CustomerID != customer.ID {
		return domain.PaymentResponse{}, errors.New("transaction forbidden")
	}

	if transaction.Status == "completed" {
		return domain.PaymentResponse{}, errors.New("transaction already paid")
	}

	err = uc.transactionRepository.UpdateStatus(transactionID, "completed")
	if err != nil {
		return domain.PaymentResponse{}, fmt.Errorf("failed to update transaction status: %w", err)
	}

	err = uc.historyRepository.CreatePaymentHistory(customer.ID, transactionID)
	if err != nil {
		return domain.PaymentResponse{}, fmt.Errorf("failed to record payment history: %w", err)
	}

	response := domain.PaymentResponse{
		TransactionID: transaction.ID,
		CustomerID:    transaction.CustomerID,
		MerchantID:    transaction.MerchantID,
		Amount:        transaction.Amount,
		Status:        "completed",
	}

	return response, nil
}
