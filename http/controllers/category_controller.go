package controllers

import (
	"MireaPR4/database/models"
	"MireaPR4/database/repositories"
	"errors"
)

type CategoryController interface {
	Create(name string) (*models.Category, error)
	GetAll() ([]models.Category, error)
	GetByID(id int) (*models.Category, error)
	Update(id int, name string) (*models.Category, error)
	Delete(id int) error
}

type categoryController struct {
	repo repositories.CategoryRepository
}

func NewCategoryController(repo repositories.CategoryRepository) CategoryController {
	return &categoryController{repo}
}

func (c *categoryController) Create(name string) (*models.Category, error) {
	// Проверка на дублирование категории
	existingCategory, err := c.repo.GetByName(name)
	if err == nil || existingCategory != nil {
		return nil, errors.New("category with this name already exists")
	}

	// Создание новой категории
	category := models.Category{
		Name: name,
	}

	// Сохраняем категорию в репозитории
	e := c.repo.Create(&category)
	if e != nil {
		return nil, e
	}

	return &category, nil
}

func (c *categoryController) GetAll() ([]models.Category, error) {
	// Получаем все категории из репозитория
	categories, err := c.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *categoryController) GetByID(id int) (*models.Category, error) {
	// Ищем категорию по ID
	category, err := c.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *categoryController) Update(id int, name string) (*models.Category, error) {
	// Получаем существующую категорию по ID
	existingCategory, err := c.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Обновляем имя категории
	existingCategory.Name = name

	// Сохраняем обновленную категорию
	err = c.repo.Update(existingCategory)
	if err != nil {
		return nil, err
	}

	return existingCategory, nil
}

func (c *categoryController) Delete(id int) error {
	// Ищем категорию по ID
	category, err := c.repo.GetByID(id)
	if err != nil {
		return err
	}

	// Удаляем категорию
	err = c.repo.Delete(category.ID)
	if err != nil {
		return err
	}

	return nil
}
