package repository

type HistoryRepository interface {
	CreateLoginHistory(customerID string) error
	CreateLogoutHistory(customerID string) error
	CreatePaymentHistory(customerID, transactionID string) error
}
