package dto

type UpdateAddressDTO struct {
	City   *string `json:"city"`
	Street *string `json:"street"`
	House  *int    `json:"house"`
	Index  *string `json:"index"`
	Flat   *int    `json:"flat"`
}
