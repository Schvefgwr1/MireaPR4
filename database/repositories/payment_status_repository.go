package repositories

import (
	"MireaPR4/database/models"
	"errors"
	"gorm.io/gorm"
)

type PaymentStatusRepository interface {
	GetByID(id int) (*models.PaymentStatus, error)
}

type paymentStatusRepository struct {
	db *gorm.DB
}

func NewPaymentStatusRepository(db *gorm.DB) PaymentStatusRepository {
	return &paymentStatusRepository{db}
}

func (r *paymentStatusRepository) GetByID(id int) (*models.PaymentStatus, error) {
	var status models.PaymentStatus
	result := r.db.First(&status, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("payment status not found")
	} else if result.Error != nil {
		return nil, result.Error
	}
	return &status, nil
}
