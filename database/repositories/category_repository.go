package repositories

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	GetAll() ([]models.Category, error)
	GetByID(id int) (*models.Category, error)
	GetByName(name string) (*models.Category, error)
	Update(category *models.Category) error
	Delete(id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetByID(id int) (*models.Category, error) {
	var category models.Category
	err := r.db.First(&category, id).Error
	return &category, err
}

func (r *categoryRepository) GetByName(name string) (*models.Category, error) {
	var category models.Category
	err := r.db.
		Where("username = ?", name).
		First(&category).Error
	return &category, err
}

func (r *categoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id int) error {
	return r.db.Delete(&models.Category{}, id).Error
}
