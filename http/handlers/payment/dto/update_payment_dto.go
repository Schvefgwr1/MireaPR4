package dto

type UpdatePaymentDTO struct {
	Amount   *float64 `json:"amount"`
	StatusID *int     `json:"status_id"`
}
