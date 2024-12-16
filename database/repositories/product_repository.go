package repositories

import (
	"MireaPR4/database/models"
	"context"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	GetAll(c context.Context) ([]models.Product, error)
	GetAllWithPagination(page, limit int, categoryID *int) ([]models.Product, error)
	GetByID(id int) (*models.Product, error)
	Count() (int64, error)
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

func (r *productRepository) GetAll(c context.Context) ([]models.Product, error) {
	var products []models.Product
	err := r.db.WithContext(c).Preload("Category").Find(&products).Error
	return products, err
}

func (r *productRepository) GetAllWithPagination(page, limit int, categoryID *int) ([]models.Product, error) {
	var products []models.Product

	query := r.db
	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	}

	offset := (page - 1) * limit
	err := query.Preload("Category").
		Limit(limit).
		Offset(offset).
		Find(&products).Error

	return products, err
}

func (r *productRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.Product{}).Count(&count).Error
	return count, err
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
