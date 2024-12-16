package dto

type CreateEmployeeDTO struct {
	UserID     int    `json:"user_id" binding:"required"`
	Position   string `json:"position" binding:"required"`
	Department string `json:"department" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	Email      string `json:"email" binding:"required"`
}
