package controllers

import (
	"MireaPR4/database/models"
	"MireaPR4/database/repositories"
	"MireaPR4/http/handlers/role/dto"
	"errors"
)

type RoleController interface {
	Create(request *dto.CreateRoleDTO) (*models.Role, error)
	GetAll() ([]models.Role, error)
	GetRoleByID(id int) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	Update(id int, request *dto.UpdateRoleDTO) (*models.Role, error)
	Delete(id int) error
}

func NewRoleController(rep repositories.RoleRepository, permRep repositories.PermissionRepository) RoleController {
	return &roleController{rep, permRep}
}

type roleController struct {
	repository           repositories.RoleRepository
	permissionRepository repositories.PermissionRepository
}

func (rc *roleController) Create(request *dto.CreateRoleDTO) (*models.Role, error) {
	var permissions []models.Permission
	for _, permId := range request.Permissions {
		permission, err := rc.permissionRepository.GetByID(permId)
		if err != nil || permission == nil {
			return nil, errors.New("Error of find permission with ID: " + string(rune(permId)))
		}
		permissions = append(permissions, *permission)
	}
	role := models.Role{
		Name:        request.RoleName,
		Permissions: permissions,
	}
	if err := rc.repository.Create(&role); err == nil {
		return &role, nil
	} else {
		return nil, err
	}
}

func (rc *roleController) GetAll() ([]models.Role, error) {
	return rc.repository.GetAll()
}

func (rc *roleController) GetRoleByID(id int) (*models.Role, error) {
	if role, err := rc.repository.GetByID(id); err != nil {
		return nil, err
	} else {
		return role, nil
	}
}

func (rc *roleController) GetRoleByName(name string) (*models.Role, error) {
	if role, err := rc.repository.GetByName(name); err != nil {
		return nil, err
	} else {
		return role, nil
	}
}

func (rc *roleController) Update(id int, request *dto.UpdateRoleDTO) (*models.Role, error) {
	var role *models.Role
	role, err := rc.repository.GetByID(id)

	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, errors.New("error of update")
	}

	var permissions []models.Permission
	for _, permId := range request.Permissions {
		permission, err := rc.permissionRepository.GetByID(permId)
		if err != nil || permission == nil {
			return nil, errors.New("Error of find permission with ID: " + string(rune(permId)))
		}
		permissions = append(permissions, *permission)
	}
	role.Permissions = permissions

	if request.Name != nil {
		role.Name = *request.Name
	}

	if err := rc.repository.Update(role); err == nil {
		return role, nil
	} else {
		return nil, err
	}
}

func (rc *roleController) Delete(id int) error {
	return rc.repository.Delete(id)
}
