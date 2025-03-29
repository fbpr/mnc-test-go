package persistent

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/fbpr/mnc-test-go/internal/domain"
	"github.com/fbpr/mnc-test-go/internal/repository"
)

type HistoryRepositoryJSON struct {
	filePath string
}

func NewHistoryRepositoryJSON(dataDir string) repository.HistoryRepository {
	return &HistoryRepositoryJSON{
		filePath: filepath.Join(dataDir, "history.json"),
	}
}

func (r *HistoryRepositoryJSON) Create(history domain.History) error {
	histories, err := r.LoadHistories()
	if err != nil {
		return err
	}

	if history.ID == "" {
		if len(histories) == 0 {
			history.ID = "1"
		} else {
			lastHistory := histories[len(histories)-1]
			id, err := strconv.Atoi(lastHistory.ID)
			if err != nil {
				history.ID = fmt.Sprintf("%d", len(histories)+1)
			} else {
				history.ID = fmt.Sprintf("%d", id+1)
			}
		}
	}

	if history.Timestamp.IsZero() {
		history.Timestamp = time.Now().UTC()
	}

	histories = append(histories, history)

	data, err := json.MarshalIndent(histories, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0o644)
}

func (r *HistoryRepositoryJSON) CreateLoginHistory(customerID string) error {
	history := domain.History{
		CustomerID: customerID,
		Action:     "login",
		Timestamp:  time.Now().UTC(),
	}
	return r.Create(history)
}

func (r *HistoryRepositoryJSON) CreateLogoutHistory(customerID string) error {
	history := domain.History{
		CustomerID: customerID,
		Action:     "logout",
		Timestamp:  time.Now().UTC(),
	}
	return r.Create(history)
}

func (r *HistoryRepositoryJSON) CreatePaymentHistory(customerID, transactionID string) error {
	history := domain.History{
		CustomerID:    customerID,
		TransactionID: transactionID,
		Action:        "payment",
		Timestamp:     time.Now().UTC(),
	}
	return r.Create(history)
}

func (r *HistoryRepositoryJSON) getNextID(histories []domain.History) string {
	if len(histories) == 0 {
		return "1"
	}

	lastHistory := histories[len(histories)-1]
	id, err := strconv.Atoi(lastHistory.ID)
	if err != nil {
		return fmt.Sprintf("%d", len(histories)+1)
	}

	return fmt.Sprintf("%d", id+1)
}

func (r *HistoryRepositoryJSON) LoadHistories() ([]domain.History, error) {
	fmt.Println(r.filePath)
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var histories []domain.History

	if len(data) == 0 {
		return []domain.History{}, nil
	}

	err = json.Unmarshal(data, &histories)
	return histories, err
}
