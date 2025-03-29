package persistent

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fbpr/mnc-test-go/internal/domain"
	"github.com/fbpr/mnc-test-go/internal/repository"
)

type TransactionRepositoryJSON struct {
	filePath string
}

func NewTransactionRepositoryJSON(dataDir string) repository.TransactionRepository {
	return &TransactionRepositoryJSON{
		filePath: filepath.Join(dataDir, "transaction.json"),
	}
}

func (r *TransactionRepositoryJSON) GetByID(id string) (domain.Transaction, error) {
	transactions, err := r.LoadTransactions()
	if err != nil {
		return domain.Transaction{}, err
	}

	for _, transaction := range transactions {
		if transaction.ID == id {
			return transaction, nil
		}
	}

	return domain.Transaction{}, errors.New("transaction not found")
}

func (r *TransactionRepositoryJSON) UpdateStatus(id, status string) error {
	transactions, err := r.LoadTransactions()
	if err != nil {
		return err
	}

	found := false
	for i, transaction := range transactions {
		if transaction.ID == id {
			transactions[i].Status = status
			found = true
			break
		}
	}

	if !found {
		return errors.New("transaction not found")
	}

	return r.SaveTransactions(transactions)
}

func (r *TransactionRepositoryJSON) LoadTransactions() ([]domain.Transaction, error) {
	fmt.Println(r.filePath)
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var transactions []domain.Transaction
	if len(data) == 0 {
		return []domain.Transaction{}, nil
	}

	err = json.Unmarshal(data, &transactions)
	return transactions, err
}

func (r *TransactionRepositoryJSON) SaveTransactions(transactions []domain.Transaction) error {
	data, err := json.MarshalIndent(transactions, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0o644)
}
