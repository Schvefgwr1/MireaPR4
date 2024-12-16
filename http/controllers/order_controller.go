package controllers

import (
	"MireaPR4/database/models"
	"MireaPR4/database/repositories"
	"MireaPR4/http/handlers/order/dto"
	"context"
	"errors"
	"strconv"
)

type OrderController interface {
	Create(request *dto.CreateOrderDTO) (*models.Order, error)
	GetAll(cont context.Context) ([]models.Order, error)
	GetAllPaginated(offset, limit int, userID *int) ([]models.Order, int64, error)
	GetByID(id int) (*models.Order, error)
	Update(id int, data any) (*models.Order, error)
	Delete(id int) error
}

type orderController struct {
	repo            repositories.OrderRepository
	userRepo        repositories.UserRepository
	orderStatusRepo repositories.OrderStatusRepository
	productRepo     repositories.ProductRepository
	orderItemRepo   repositories.OrderItemRepository
}

func NewOrderController(
	repo repositories.OrderRepository,
	userRepo repositories.UserRepository,
	orderStatusRepo repositories.OrderStatusRepository,
	productRepo repositories.ProductRepository,
	orderItemRepo repositories.OrderItemRepository,
) OrderController {
	return &orderController{
		repo,
		userRepo,
		orderStatusRepo,
		productRepo,
		orderItemRepo,
	}
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

func (c *orderController) GetAll(cont context.Context) ([]models.Order, error) {
	return c.repo.GetAll(cont)
}

func (c *orderController) GetAllPaginated(offset, limit int, userID *int) ([]models.Order, int64, error) {
	orders, err := c.repo.GetAllPaginated(offset, limit, userID)
	if err != nil {
		return nil, 0, err
	}

	count, err := c.repo.Count()
	if err != nil {
		return nil, 0, err
	}

	return orders, count, nil
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
	order, err := c.repo.GetByID(id)
	if err != nil {
		return errors.New("error of orders repo")
	}

	for _, item := range order.OrderItems {
		if err := c.orderItemRepo.Delete(item.ID); err != nil {
			return errors.New("error of items repo")
		}
	}

	return c.repo.Delete(id)
}
