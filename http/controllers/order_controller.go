package controllers

import (
	"MireaPR4/database/models"
	"MireaPR4/database/repositories"
	"MireaPR4/http/handlers/order/dto"
	"errors"
	"strconv"
)

type OrderController interface {
	Create(request *dto.CreateOrderDTO) (*models.Order, error)
	GetAll() ([]models.Order, error)
	GetByID(id int) (*models.Order, error)
	Update(id int, data any) (*models.Order, error)
	Delete(id int) error
}

type orderController struct {
	repo            repositories.OrderRepository
	userRepo        repositories.UserRepository
	orderStatusRepo repositories.OrderStatusRepository
	productRepo     repositories.ProductRepository
}

func NewOrderController(
	repo repositories.OrderRepository,
	userRepo repositories.UserRepository,
	orderStatusRepo repositories.OrderStatusRepository,
	productRepo repositories.ProductRepository,
) OrderController {
	return &orderController{repo, userRepo, orderStatusRepo, productRepo}
}

func (c *orderController) Create(request *dto.CreateOrderDTO) (*models.Order, error) {
	if val, err := c.userRepo.GetByID(request.UserID); err != nil || val == nil {
		return nil, errors.New("Incorrect UserID: " + strconv.Itoa(request.UserID))
	}

	if val, err := c.orderStatusRepo.GetByID(request.StatusID); err != nil || val == nil {
		return nil, errors.New("Incorrect StatusID: " + strconv.Itoa(request.StatusID))
	}

	totalPrice := 0.0
	var orderItems []*models.OrderItem

	for _, item := range request.Items {
		product, err := c.productRepo.GetByID(item.ProductID)
		if err != nil || product == nil {
			return nil, errors.New("Incorrect ProductID: " + strconv.Itoa(item.ProductID))
		}

		orderItem := models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}

		if item.Price != nil {
			orderItem.Price = *item.Price
		} else {
			orderItem.Price = product.Price
		}

		totalPrice += orderItem.Price * float64(orderItem.Quantity)
		orderItems = append(orderItems, &orderItem)
	}

	order := models.Order{
		UserID:     request.UserID,
		StatusID:   request.StatusID,
		TotalPrice: totalPrice,
		OrderItems: orderItems,
	}

	if err := c.repo.Create(&order); err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *orderController) GetAll() ([]models.Order, error) {
	return c.repo.GetAll()
}

func (c *orderController) GetByID(id int) (*models.Order, error) {
	return c.repo.GetByID(id)
}

func (c *orderController) Update(id int, data any) (*models.Order, error) {
	order := data.(*models.Order)
	order.ID = id
	if err := c.repo.Update(order); err != nil {
		return nil, err
	}
	return order, nil
}

func (c *orderController) Delete(id int) error {
	return c.repo.Delete(id)
}
