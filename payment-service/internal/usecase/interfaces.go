package usecase

import "payment-service/internal/domain"

type PaymentRepository interface {
	Save(payment *domain.Payment) error
	GetByOrderID(orderID string) (*domain.Payment, error)
}
