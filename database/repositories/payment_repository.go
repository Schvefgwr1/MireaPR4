package repositories

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *models.Payment) (*models.Payment, error)
	GetAll() ([]models.Payment, error)
	GetByID(id int) (*models.Payment, error)
	Update(payment *models.Payment) (*models.Payment, error)
	Delete(id int) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) Create(payment *models.Payment) (*models.Payment, error) {
	if err := r.db.Create(payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *paymentRepository) GetAll() ([]models.Payment, error) {
	var payments []models.Payment
	if err := r.db.Preload("Status").Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *paymentRepository) GetByID(id int) (*models.Payment, error) {
	var payment models.Payment
	if err := r.db.Preload("Status").First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) Update(payment *models.Payment) (*models.Payment, error) {
	if err := r.db.Updates(payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *paymentRepository) Delete(id int) error {
	return r.db.Delete(&models.Payment{}, id).Error
}
