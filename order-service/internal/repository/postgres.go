package repository

import (
	"order-service/internal/domain"

	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepo {
	return &orderRepo{db: db}
}

func (r *orderRepo) Save(order *domain.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepo) GetByID(id string) (*domain.Order, error) {
	var order domain.Order
	err := r.db.First(&order, "id = ?", id).Error
	return &order, err
}

func (r *orderRepo) Update(order *domain.Order) error {
	return r.db.Save(order).Error
}
func (r *orderRepo) GetByIdempotencyKey(key string) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Where("idempotency_key = ?", key).First(&order).Error
	return &order, err
}
