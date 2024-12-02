package repositories

import (
	"MireaPR4/database/models"
	"errors"
	"gorm.io/gorm"
)

type ShipmentStatusRepository interface {
	Create(status *models.ShipmentStatus) (*models.ShipmentStatus, error)
	GetAll() ([]models.ShipmentStatus, error)
	GetByID(id int) (*models.ShipmentStatus, error)
	Update(id int, updatedStatus *models.ShipmentStatus) (*models.ShipmentStatus, error)
	Delete(id int) error
}

type shipmentStatusRepository struct {
	db *gorm.DB
}

func NewShipmentStatusRepository(db *gorm.DB) ShipmentStatusRepository {
	return &shipmentStatusRepository{db}
}

func (r *shipmentStatusRepository) Create(status *models.ShipmentStatus) (*models.ShipmentStatus, error) {
	result := r.db.Create(status)
	if result.Error != nil {
		return nil, result.Error
	}
	return status, nil
}

func (r *shipmentStatusRepository) GetAll() ([]models.ShipmentStatus, error) {
	var statuses []models.ShipmentStatus
	result := r.db.Find(&statuses)
	if result.Error != nil {
		return nil, result.Error
	}
	return statuses, nil
}

func (r *shipmentStatusRepository) GetByID(id int) (*models.ShipmentStatus, error) {
	var status models.ShipmentStatus
	result := r.db.First(&status, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("shipment status not found")
	} else if result.Error != nil {
		return nil, result.Error
	}
	return &status, nil
}

func (r *shipmentStatusRepository) Update(id int, updatedStatus *models.ShipmentStatus) (*models.ShipmentStatus, error) {
	status, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}

	status.Name = updatedStatus.Name

	result := r.db.Save(status)
	if result.Error != nil {
		return nil, result.Error
	}
	return status, nil
}

func (r *shipmentStatusRepository) Delete(id int) error {
	status, err := r.GetByID(id)
	if err != nil {
		return err
	}

	result := r.db.Delete(&status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
