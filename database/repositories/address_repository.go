package repositories

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
)

type AddressRepository interface {
	Create(address *models.Address) error
	GetAll() ([]models.Address, error)
	GetByID(id int) (*models.Address, error)
	Update(address *models.Address) error
	Delete(id int) error
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db}
}

func (r *addressRepository) Create(address *models.Address) error {
	return r.db.Create(address).Error
}

func (r *addressRepository) GetAll() ([]models.Address, error) {
	var addresses []models.Address
	err := r.db.Find(&addresses).Error
	return addresses, err
}

func (r *addressRepository) GetByID(id int) (*models.Address, error) {
	var address models.Address
	err := r.db.First(&address, id).Error
	return &address, err
}

func (r *addressRepository) Update(address *models.Address) error {
	return r.db.Save(address).Error
}

func (r *addressRepository) Delete(id int) error {
	return r.db.Delete(&models.Address{}, id).Error
}
