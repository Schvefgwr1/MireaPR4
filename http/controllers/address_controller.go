package controllers

import (
	"MireaPR4/database/models"
	"MireaPR4/database/repositories"
	"MireaPR4/http/handlers/address/dto"
)

type AddressController interface {
	Create(request *dto.CreateAddressDTO) (*models.Address, error)
	GetAll() ([]models.Address, error)
	GetByID(id int) (*models.Address, error)
	Update(id int, request *dto.UpdateAddressDTO) (*models.Address, error)
	Delete(id int) error
}

type addressController struct {
	repo repositories.AddressRepository
}

func NewAddressController(repo repositories.AddressRepository) AddressController {
	return &addressController{repo}
}

func (c *addressController) Create(request *dto.CreateAddressDTO) (*models.Address, error) {
	address := models.Address{
		City:   request.City,
		Street: request.Street,
		House:  request.House,
		Index:  request.Index,
		Flat:   request.Flat,
	}

	if err := c.repo.Create(&address); err != nil {
		return nil, err
	}

	return &address, nil
}

func (c *addressController) GetAll() ([]models.Address, error) {
	return c.repo.GetAll()
}

func (c *addressController) GetByID(id int) (*models.Address, error) {
	return c.repo.GetByID(id)
}

func (c *addressController) Update(id int, request *dto.UpdateAddressDTO) (*models.Address, error) {
	address, err := c.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if request.City != nil {
		address.City = *request.City
	}
	if request.Street != nil {
		address.Street = *request.Street
	}
	if request.House != nil {
		address.House = *request.House
	}
	if request.Index != nil {
		address.Index = *request.Index
	}
	if request.Flat != nil {
		address.Flat = *request.Flat
	}

	if err := c.repo.Update(address); err != nil {
		return nil, err
	}

	return address, nil
}

func (c *addressController) Delete(id int) error {
	return c.repo.Delete(id)
}
