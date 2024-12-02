package controllers

import (
	"MireaPR4/database/models"
	"MireaPR4/database/repositories"
	"MireaPR4/http/handlers/payment/dto"
	"errors"
)

type PaymentController interface {
	Create(data *dto.CreatePaymentDTO) (*models.Payment, error)
	GetAll() ([]models.Payment, error)
	GetByID(id int) (*models.Payment, error)
	Update(id int, data *dto.UpdatePaymentDTO) (*models.Payment, error)
	Delete(id int) error
}

type paymentController struct {
	repo repositories.PaymentRepository
}

func NewPaymentController(repo repositories.PaymentRepository) PaymentController {
	return &paymentController{repo}
}

func (c *paymentController) Create(data *dto.CreatePaymentDTO) (*models.Payment, error) {
	payment := &models.Payment{
		OrderID: data.OrderID,
		Amount:  data.Amount,
	}

	createdPayment, err := c.repo.Create(payment)
	if err != nil {
		return nil, err
	}

	return createdPayment, nil
}

func (c *paymentController) GetAll() ([]models.Payment, error) {
	return c.repo.GetAll()
}

func (c *paymentController) GetByID(id int) (*models.Payment, error) {
	return c.repo.GetByID(id)
}

func (c *paymentController) Update(id int, data *dto.UpdatePaymentDTO) (*models.Payment, error) {
	payment, err := c.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("payment not found")
	}

	// Обновление статуса, если передан
	if data.StatusID != nil {
		payment.StatusID = *data.StatusID
	}

	// Обновление суммы, если передана
	if data.Amount != nil {
		payment.Amount = *data.Amount
	}

	updatedPayment, err := c.repo.Update(payment)
	if err != nil {
		return nil, err
	}

	return updatedPayment, nil
}

func (c *paymentController) Delete(id int) error {
	return c.repo.Delete(id)
}
