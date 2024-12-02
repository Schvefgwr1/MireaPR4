package repositories

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
)

type PermissionRepository interface {
	GetAll() ([]models.Permission, error)
	GetByID(id int) (*models.Permission, error)
	GetByName(name string) (*models.Permission, error)
}

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{db}
}

func (r *permissionRepository) GetAll() ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.db.Find(&permissions).Error
	return permissions, err
}

func (r *permissionRepository) GetByID(id int) (*models.Permission, error) {
	var permission models.Permission
	err := r.db.First(&permission, id).Error
	return &permission, err
}

func (r *permissionRepository) GetByName(username string) (*models.Permission, error) {
	var permission models.Permission
	err := r.db.
		Where("username = ?", username).
		First(&permission).Error
	return &permission, err
}
