package domain

import "time"

type History struct {
	ID            string    `json:"id"`
	CustomerID    string    `json:"customer_id"`
	TransactionID string    `json:"transaction_id,omitempty"`
	Action        string    `json:"action"`
	Timestamp     time.Time `json:"timestamp"`
}
