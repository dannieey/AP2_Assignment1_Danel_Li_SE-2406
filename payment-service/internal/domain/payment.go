package domain

type Payment struct {
	ID            string `json:"id" gorm:"primaryKey"`
	OrderID       string `json:"order_id"`
	TransactionID string `json:"transaction_id"`
	Amount        int64  `json:"amount"`
	Status        string `json:"status"`
}
