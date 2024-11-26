package dto

type CreateOrderItemDto struct {
	ProductID int      `json:"product_id" binding:"required"`
	Quantity  int      `json:"quantity" binding:"required"`
	Price     *float64 `json:"price"`
}
