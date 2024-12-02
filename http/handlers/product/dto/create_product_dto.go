package dto

type CreateProductDTO struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,min=0"`
	Stock       int     `json:"stock" binding:"required"`
	CategoryID  int     `json:"category_id" binding:"required"`
}
