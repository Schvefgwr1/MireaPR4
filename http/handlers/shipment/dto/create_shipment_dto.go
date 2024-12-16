package dto

type CreateShipmentDTO struct {
	OrderID   int `json:"order_id" binding:"required"`
	StatusID  int `json:"status_id" binding:"required"`
	AddressId int `json:"address" binding:"required"`
}
