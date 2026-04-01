package usecase

import "order-service/internal/domain"

type OrderRepository interface {
	Save(order *domain.Order) error
	GetByID(id string) (*domain.Order, error)
	Update(order *domain.Order) error
	GetByIdempotencyKey(key string) (*domain.Order, error)
}
type PaymentClient interface {
	CheckPayment(orderID string, amount int64) (string, error)
}
