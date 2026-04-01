package repository

import (
	"payment-service/internal/domain"

	"gorm.io/gorm"
)

type postgresRepo struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *postgresRepo {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) Save(p *domain.Payment) error {
	return r.db.Create(p).Error
}

func (r *postgresRepo) GetByOrderID(orderID string) (*domain.Payment, error) {
	var p domain.Payment
	err := r.db.Where("order_id = ?", orderID).First(&p).Error
	return &p, err
}
