package repositories

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
)

type ShipmentRepository interface {
	Create(shipment *models.Shipment) error
	GetAll() ([]models.Shipment, error)
	GetByID(id int) (*models.Shipment, error)
	Update(shipment *models.Shipment) error
	Delete(id int) error
}

type shipmentRepository struct {
	db *gorm.DB
}

func NewShipmentRepository(db *gorm.DB) ShipmentRepository {
	return &shipmentRepository{db}
}

func (r *shipmentRepository) Create(shipment *models.Shipment) error {
	return r.db.Create(shipment).Error
}

func (r *shipmentRepository) GetAll() ([]models.Shipment, error) {
	var shipments []models.Shipment
	err := r.db.Find(&shipments).Error
	return shipments, err
}

func (r *shipmentRepository) GetByID(id int) (*models.Shipment, error) {
	var shipment models.Shipment
	err := r.db.First(&shipment, id).Error
	return &shipment, err
}

func (r *shipmentRepository) Update(shipment *models.Shipment) error {
	return r.db.Save(shipment).Error
}

func (r *shipmentRepository) Delete(id int) error {
	return r.db.Delete(&models.Shipment{}, id).Error
}
