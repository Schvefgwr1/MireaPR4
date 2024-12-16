package dto

type CreateOrderDTO struct {
	UserID   int                  `json:"user_id"`
	StatusID int                  `json:"status_id"`
	Items    []CreateOrderItemDto `json:"items"`
}
