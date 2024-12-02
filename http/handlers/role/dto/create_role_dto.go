package dto

type CreateRoleDTO struct {
	RoleName    string `json:"name"`
	Permissions []int  `json:"permissions"`
}
