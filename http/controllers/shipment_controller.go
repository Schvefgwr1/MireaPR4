package controllers

import (
	"MireaPR4/database/models"
	"MireaPR4/database/repositories"
	"MireaPR4/http/handlers/shipment/dto"
	"errors"
)

type ShipmentController interface {
	Create(request *dto.CreateShipmentDTO) (*models.Shipment, error)
	GetAll() ([]models.Shipment, error)
	GetByID(id int) (*models.Shipment, error)
	Update(id int, data *dto.UpdateShipmentDTO) (*models.Shipment, error)
	Delete(id int) error
}

type shipmentController struct {
	repo               repositories.ShipmentRepository
	orderRepo          repositories.OrderRepository
	shipmentStatusRepo repositories.ShipmentStatusRepository
}

func NewShipmentController(repo repositories.ShipmentRepository, orderRepo repositories.OrderRepository, shipmentStatusRepo repositories.ShipmentStatusRepository) ShipmentController {
	return &shipmentController{repo, orderRepo, shipmentStatusRepo}
}

func (c *shipmentController) Create(request *dto.CreateShipmentDTO) (*models.Shipment, error) {
	order, err := c.orderRepo.GetByID(request.OrderID)
	if err != nil || order == nil {
		return nil, errors.New("invalid Order ID")
	}

	status, err := c.shipmentStatusRepo.GetByID(request.StatusID)
	if err != nil || status == nil {
		return nil, errors.New("invalid Status ID")
	}

	shipment := models.Shipment{
		OrderID:   request.OrderID,
		StatusID:  request.StatusID,
		AddressID: request.AddressId,
	}

	if err := c.repo.Create(&shipment); err != nil {
		return nil, err
	}

	return &shipment, nil
}

func (c *shipmentController) GetAll() ([]models.Shipment, error) {
	return c.repo.GetAll()
}

func (c *shipmentController) GetByID(id int) (*models.Shipment, error) {
	return c.repo.GetByID(id)
}

func (c *shipmentController) Update(id int, data *dto.UpdateShipmentDTO) (*models.Shipment, error) {
	shipment, err := c.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if data.StatusID != nil {
		shipment.StatusID = *data.StatusID
	}
	if data.AddressID != nil {
		shipment.AddressID = *data.AddressID
	}

	if err := c.repo.Update(shipment); err != nil {
		return nil, err
	}

	return shipment, nil
}

func (c *shipmentController) Delete(id int) error {
	return c.repo.Delete(id)
}
