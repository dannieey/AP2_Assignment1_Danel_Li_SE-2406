package usecase

import (
	"payment-service/internal/domain"

	"github.com/google/uuid"
)

type PaymentUseCase struct {
	repo PaymentRepository
}

func NewPaymentUseCase(r PaymentRepository) *PaymentUseCase {
	return &PaymentUseCase{repo: r}
}

func (uc *PaymentUseCase) ProcessPayment(orderID string, amount int64) (*domain.Payment, error) {
	status := "Authorized"
	if amount > 100000 {
		status = "Declined"
	}

	payment := &domain.Payment{
		ID:            uuid.New().String(),
		OrderID:       orderID,
		TransactionID: uuid.New().String(),
		Amount:        amount,
		Status:        status,
	}

	if err := uc.repo.Save(payment); err != nil {
		return nil, err
	}

	return payment, nil
}

func (uc *PaymentUseCase) GetPaymentStatus(orderID string) (*domain.Payment, error) {
	return uc.repo.GetByOrderID(orderID)
}
