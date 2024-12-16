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
	repo       repositories.PaymentRepository
	statusRepo repositories.PaymentStatusRepository
}

func NewPaymentController(
	repo repositories.PaymentRepository,
	statusRepo repositories.PaymentStatusRepository,
) PaymentController {
	return &paymentController{repo, statusRepo}
}

func (c *paymentController) Create(data *dto.CreatePaymentDTO) (*models.Payment, error) {
	payment := &models.Payment{
		OrderID:  data.OrderID,
		Amount:   data.Amount,
		StatusID: 1,
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
		status, e := c.statusRepo.GetByID(*data.StatusID)
		if e != nil || status == nil {
			return nil, errors.New("payment status not found")
		}
		payment.StatusID = *data.StatusID
		payment.Status = *status
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
