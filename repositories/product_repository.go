package repositories

import (
	"MireaPR4/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	GetAll() ([]models.Product, error)
	GetByID(id int) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id int) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Category").Find(&products).Error
	return products, err
}

func (r *productRepository) GetByID(id int) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Category").First(&product, id).Error
	return &product, err
}

func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id int) error {
	return r.db.Delete(&models.Product{}, id).Error
}
