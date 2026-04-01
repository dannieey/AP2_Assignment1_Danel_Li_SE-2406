package usecase

import (
	"errors"
	"order-service/internal/domain"
	"time"

	"github.com/google/uuid"
)

type OrderUseCase struct {
	repo          OrderRepository
	paymentClient PaymentClient
}

func NewOrderUseCase(r OrderRepository, pc PaymentClient) *OrderUseCase {
	return &OrderUseCase{repo: r, paymentClient: pc}
}

func (uc *OrderUseCase) CreateOrder(itemName string, customerID string, amount int64, idempKey string) (*domain.Order, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	existingOrder, err := uc.repo.GetByIdempotencyKey(idempKey)
	if err == nil {
		return existingOrder, nil
	}

	order := &domain.Order{
		ID:             uuid.New().String(),
		IdempotencyKey: idempKey,
		CustomerID:     customerID,
		ItemName:       itemName,
		Amount:         amount,
		Status:         "Pending",
		CreatedAt:      time.Now(),
	}

	if err := uc.repo.Save(order); err != nil {
		return nil, err
	}

	status, err := uc.paymentClient.CheckPayment(order.ID, order.Amount)

	if err != nil {
		return nil, errors.New("Payment Service is unavailable, please try again later")
	}

	if status == "Declined" {
		order.Status = "Failed"
	} else {
		order.Status = "Paid"
	}

	err = uc.repo.Update(order)
	return order, err
}

func (uc *OrderUseCase) CancelOrder(orderID string) error {
	order, err := uc.repo.GetByID(orderID)
	if err != nil {
		return err
	}

	if order.Status == "Paid" {
		return errors.New("cannot cancel a paid order")
	}

	order.Status = "Cancelled"
	return uc.repo.Update(order)
}

func (uc *OrderUseCase) GetOrder(id string) (*domain.Order, error) {
	order, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}
