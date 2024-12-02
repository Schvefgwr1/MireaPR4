package repositories

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role *models.Role) error
	GetAll() ([]models.Role, error)
	GetByID(id int) (*models.Role, error)
	GetByName(name string) (*models.Role, error)
	Update(user *models.Role) error
	Delete(id int) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) Create(role *models.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) GetAll() ([]models.Role, error) {
	var roles []models.Role
	err := r.db.Preload("Permissions").Find(&roles).Error
	return roles, err
}

func (r *roleRepository) GetByID(id int) (*models.Role, error) {
	var role models.Role
	err := r.db.
		Preload("Permissions").
		First(&role, id).Error
	return &role, err
}

func (r *roleRepository) GetByName(username string) (*models.Role, error) {
	var role models.Role
	err := r.db.
		Preload("Permissions").
		Where("username = ?", username).
		First(&role).Error
	return &role, err
}

func (r *roleRepository) Update(role *models.Role) error {
	return r.db.Save(role).Error
}

func (r *roleRepository) Delete(id int) error {
	return r.db.Delete(&models.Role{}, id).Error
}
