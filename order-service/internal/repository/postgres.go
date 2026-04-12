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

func (r *orderRepo) GetStats() (map[string]int64, error) {
	var results []struct {
		Status string
		Count  int64
	}

	err := r.db.Model(&domain.Order{}).
		Select("status, count(*) as count").
		Group("status").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	statsMap := make(map[string]int64)
	for _, res := range results {
		statsMap[res.Status] = res.Count
	}

	return statsMap, nil
}
