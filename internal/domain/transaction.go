package domain

type Transaction struct {
	ID         string  `json:"id"`
	CustomerID string  `json:"customer_id"`
	MerchantID string  `json:"merchant_id"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"`
	Timestamp  string  `json:"timestamp"`
}

type PaymentRequest struct {
	CustomerID    string `json:"customer_id"`
	CustomerEmail string `json:"customer_email"`
}

type PaymentResponse struct {
	TransactionID string  `json:"transaction_id"`
	CustomerID    string  `json:"customer_id"`
	MerchantID    string  `json:"merchant_id"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"`
}
