package dto

type CreatePaymentDTO struct {
	OrderID int     `json:"order_id" binding:"required"`
	Amount  float64 `json:"amount" binding:"required"`
}
