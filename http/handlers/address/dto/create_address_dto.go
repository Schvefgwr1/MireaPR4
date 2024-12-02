package dto

type CreateAddressDTO struct {
	City   string `json:"city" binding:"required"`
	Street string `json:"street" binding:"required"`
	House  int    `json:"house" binding:"required"`
	Index  string `json:"index" binding:"required"`
	Flat   int    `json:"flat" binding:"required"`
}
