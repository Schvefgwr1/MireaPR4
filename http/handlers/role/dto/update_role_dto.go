package dto

type UpdateRoleDTO struct {
	Name        *string `json:"name"`
	Permissions []int   `json:"permissions" binding:"required"`
}
