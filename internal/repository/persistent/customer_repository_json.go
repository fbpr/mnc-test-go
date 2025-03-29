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

type CustomerRepositoryJSON struct {
	filePath string
}

func NewCustomerRepositoryJSON(dataDir string) repository.CustomerRepository {
	return &CustomerRepositoryJSON{
		filePath: filepath.Join(dataDir, "customer.json"),
	}
}

func (r *CustomerRepositoryJSON) GetByEmail(email string) (*domain.Customer, error) {
	customers, err := r.LoadCustomers()
	if err != nil {
		return nil, err
	}

	for _, customer := range customers {
		if customer.Email == email {
			return &customer, nil
		}
	}

	return nil, errors.New("customer not found")
}

func (r *CustomerRepositoryJSON) UpdateLoginStatus(email string, isLoggedIn bool) error {
	customers, err := r.LoadCustomers()
	if err != nil {
		return err
	}

	found := false
	for i, customer := range customers {
		if customer.Email == email {
			customers[i].Status.IsLoggedIn = isLoggedIn
			found = true
			break
		}
	}

	if !found {
		return errors.New("customer not found")
	}

	data, err := json.MarshalIndent(customers, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0o644)
}

func (r *CustomerRepositoryJSON) LoadCustomers() ([]domain.Customer, error) {
	fmt.Println(r.filePath)
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var customers []domain.Customer

	if len(data) == 0 {
		return []domain.Customer{}, nil
	}

	err = json.Unmarshal(data, &customers)
	return customers, err
}
