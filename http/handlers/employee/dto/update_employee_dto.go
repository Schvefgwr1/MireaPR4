package dto

type UpdateEmployeeDTO struct {
	Position   *string `json:"position,omitempty"`
	Department *string `json:"department,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	Email      *string `json:"email,omitempty"`
}
