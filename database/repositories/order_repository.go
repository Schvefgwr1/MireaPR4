package repositories

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *models.Order) error
	GetAll() ([]models.Order, error)
	GetByID(id int) (*models.Order, error)
	Update(order *models.Order) error
	Delete(id int) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) GetAll() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.
		Preload("OrderItems").
		Preload("Shipments").
		Preload("Payments").
		Preload("Status").
		Find(&orders).Error
	return orders, err
}

func (r *orderRepository) GetByID(id int) (*models.Order, error) {
	var order models.Order
	err := r.db.
		Preload("OrderItems").
		Preload("Shipments").
		Preload("Payments").
		Preload("Status").
		First(&order, id).Error
	return &order, err
}

func (r *orderRepository) Update(order *models.Order) error {
	return r.db.Save(order).Error
}

func (r *orderRepository) Delete(id int) error {
	return r.db.Delete(&models.Order{}, id).Error
}
