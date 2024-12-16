package repositories

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
)

type OrderStatusRepository interface {
	Create(orderSt *models.OrderStatus) error
	GetAll() ([]models.OrderStatus, error)
	GetByID(id int) (*models.OrderStatus, error)
	Update(orderSt *models.OrderStatus) error
	Delete(id int) error
}

type orderStatusRepository struct {
	db *gorm.DB
}

func NewOrderStatusRepository(db *gorm.DB) OrderStatusRepository {
	return &orderStatusRepository{db}
}

func (r *orderStatusRepository) Create(orderStatus *models.OrderStatus) error {
	return r.db.Create(orderStatus).Error
}

func (r *orderStatusRepository) GetAll() ([]models.OrderStatus, error) {
	var orderSt []models.OrderStatus
	err := r.db.Find(&orderSt).Error
	return orderSt, err
}

func (r *orderStatusRepository) GetByID(id int) (*models.OrderStatus, error) {
	var orderSt models.OrderStatus
	err := r.db.First(&orderSt, id).Error
	return &orderSt, err
}

func (r *orderStatusRepository) Update(orderSt *models.OrderStatus) error {
	return r.db.Save(orderSt).Error
}

func (r *orderStatusRepository) Delete(id int) error {
	return r.db.Delete(&models.OrderStatus{}, id).Error
}
