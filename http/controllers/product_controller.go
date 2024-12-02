package controllers

import (
	"MireaPR4/database/models"
	"MireaPR4/database/repositories"
	"MireaPR4/http/handlers/product/dto"
	"context"
	"errors"
)

type ProductController interface {
	Create(request *dto.CreateProductDTO) (*models.Product, error)
	GetAll(c *context.Context) ([]models.Product, error)
	GetAllWithPagination(page, limit int, categoryID *int) ([]models.Product, error)
	GetByID(id int) (*models.Product, error)
	Update(id int, data *dto.UpdateProductDTO) (*models.Product, error)
	Delete(id int) error
}

type productController struct {
	repo         repositories.ProductRepository
	categoryRepo repositories.CategoryRepository
}

func NewProductController(repo repositories.ProductRepository, categoryRepo repositories.CategoryRepository) ProductController {
	return &productController{repo, categoryRepo}
}

func (c *productController) Create(request *dto.CreateProductDTO) (*models.Product, error) {
	category, err := c.categoryRepo.GetByID(request.CategoryID)
	if err != nil || category == nil {
		return nil, errors.New("invalid Category ID")
	}

	product := models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		CategoryID:  request.CategoryID,
	}

	if err := c.repo.Create(&product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (c *productController) GetAll(cont *context.Context) ([]models.Product, error) {
	return c.repo.GetAll(cont)
}

func (c *productController) GetAllWithPagination(page, limit int, categoryID *int) ([]models.Product, error) {
	return c.repo.GetAllWithPagination(page, limit, categoryID)
}

func (c *productController) GetByID(id int) (*models.Product, error) {
	return c.repo.GetByID(id)
}

func (c *productController) Update(id int, data *dto.UpdateProductDTO) (*models.Product, error) {
	product, err := c.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if data.Name != nil {
		product.Name = *data.Name
	}
	if data.Description != nil {
		product.Description = *data.Description
	}
	if data.Price != nil {
		product.Price = *data.Price
	}
	if data.Stock != nil {
		product.Price = *data.Price
	}
	if data.CategoryID != nil {
		product.CategoryID = *data.CategoryID
	}

	if err := c.repo.Update(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (c *productController) Delete(id int) error {
	return c.repo.Delete(id)
}
