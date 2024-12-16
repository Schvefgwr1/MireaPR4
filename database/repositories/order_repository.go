package repositories

import (
	"MireaPR4/database/models"
	"context"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *models.Order) error
	GetAll(c context.Context) ([]models.Order, error)
	GetAllPaginated(offset, limit int, userID *int) ([]models.Order, error)
	Count() (int64, error)
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
	// Начало транзакции
	tx := r.db.Begin()

	// Создание заказа
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Подгрузка связанных данных
	if err := tx.Preload("Payments").Preload("OrderItems.Product").First(order, order.ID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Подтверждение транзакции
	return tx.Commit().Error
}

func (r *orderRepository) GetAll(c context.Context) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.
		WithContext(c).
		Preload("OrderItems").
		Preload("Shipments").
		Preload("Payments").
		Preload("Status").
		Find(&orders).Error
	return orders, err
}

func (r *orderRepository) GetAllPaginated(offset, limit int, userID *int) ([]models.Order, error) {
	var orders []models.Order

	query := r.db

	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}

	err := query.
		Preload("OrderItems").
		Preload("Shipments").
		Preload("Payments").
		Preload("Status").
		Offset(offset).
		Limit(limit).
		Find(&orders).Error
	return orders, err
}

func (r *orderRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.Order{}).Count(&count).Error
	return count, err
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
