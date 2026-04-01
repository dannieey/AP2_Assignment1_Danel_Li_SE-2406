package domain

import "time"

type Order struct {
	ID             string    `json:"id" gorm:"primaryKey"`
	IdempotencyKey string    `json:"idempotency_key" gorm:"uniqueIndex"`
	CustomerID     string    `json:"customer_id"`
	ItemName       string    `json:"item_name"`
	Amount         int64     `json:"amount"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}
