package controllers

import (
	"MireaPR4/database/models"
	"MireaPR4/database/repositories"
	"MireaPR4/http/handlers/employee/dto"
	"errors"
)

type EmployeeController interface {
	Create(employeeDTO *dto.CreateEmployeeDTO) (*models.Employee, error)
	GetAll() ([]models.Employee, error)
	GetByID(id int) (*models.Employee, error)
	Update(id int, employeeDTO *dto.UpdateEmployeeDTO) (*models.Employee, error)
	Delete(id int) error
}

type employeeController struct {
	repo repositories.EmployeeRepository
}

func NewEmployeeController(repo repositories.EmployeeRepository) EmployeeController {
	return &employeeController{repo}
}

func (c *employeeController) Create(employeeDTO *dto.CreateEmployeeDTO) (*models.Employee, error) {
	// Проверяем, существует ли пользователь
	user, err := c.repo.FindUserByID(employeeDTO.UserID)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	oldEmployee, e := c.repo.FindEmployeeByUserID(employeeDTO.UserID)
	if e == nil && oldEmployee != nil {
		return nil, errors.New("employee with userID already exists")
	}

	// Создаем сотрудника
	employee := &models.Employee{
		UserID:     employeeDTO.UserID,
		Position:   employeeDTO.Position,
		Department: employeeDTO.Department,
		Phone:      employeeDTO.Phone,
		Email:      employeeDTO.Email,
	}
	return c.repo.Create(employee)
}

func (c *employeeController) GetAll() ([]models.Employee, error) {
	return c.repo.GetAll()
}

func (c *employeeController) GetByID(id int) (*models.Employee, error) {
	if emp, err := c.repo.GetByID(id); err != nil || emp == nil {
		return nil, errors.New("employee not found")
	} else {
		return emp, nil
	}
}

func (c *employeeController) Update(id int, employeeDTO *dto.UpdateEmployeeDTO) (*models.Employee, error) {
	// Найти существующего сотрудника
	employee, err := c.repo.GetByID(id)
	if err != nil || employee == nil {
		return nil, errors.New("employee not found")
	}

	// Обновляем только те поля, которые переданы в DTO
	if employeeDTO.Position != nil {
		employee.Position = *employeeDTO.Position
	}
	if employeeDTO.Department != nil {
		employee.Department = *employeeDTO.Department
	}
	if employeeDTO.Phone != nil {
		employee.Phone = *employeeDTO.Phone
	}
	if employeeDTO.Email != nil {
		employee.Email = *employeeDTO.Email
	}

	// Сохранение изменений в базе данных
	updatedEmployee, err := c.repo.Update(id, employee)
	if err != nil {
		return nil, err
	}

	return updatedEmployee, nil
}

func (c *employeeController) Delete(id int) error {
	return c.repo.Delete(id)
}
