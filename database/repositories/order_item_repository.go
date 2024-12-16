package repositories

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	Delete(id int) error
}

type orderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepository{db}
}

func (r *orderItemRepository) Delete(id int) error {
	return r.db.Delete(&models.OrderItem{}, id).Error
}
