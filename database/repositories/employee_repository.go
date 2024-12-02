package repositories

import (
	"MireaPR4/database/models"
	"errors"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(employee *models.Employee) (*models.Employee, error)
	GetAll() ([]models.Employee, error)
	GetByID(id int) (*models.Employee, error)
	Update(id int, employee *models.Employee) (*models.Employee, error)
	Delete(id int) error
	FindUserByID(userID int) (*models.User, error)
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db}
}

func (r *employeeRepository) Create(employee *models.Employee) (*models.Employee, error) {
	result := r.db.Create(employee)
	if result.Error != nil {
		return nil, result.Error
	}
	return employee, nil
}

func (r *employeeRepository) GetAll() ([]models.Employee, error) {
	var employees []models.Employee
	result := r.db.Preload("User").Find(&employees)
	return employees, result.Error
}

func (r *employeeRepository) GetByID(id int) (*models.Employee, error) {
	var employee models.Employee
	result := r.db.Preload("User").First(&employee, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &employee, result.Error
}

func (r *employeeRepository) Update(id int, employee *models.Employee) (*models.Employee, error) {
	existing, err := r.GetByID(id)
	if err != nil || existing == nil {
		return nil, errors.New("employee not found")
	}

	result := r.db.Model(existing).Updates(employee)
	if result.Error != nil {
		return nil, result.Error
	}
	return existing, nil
}

func (r *employeeRepository) Delete(id int) error {
	result := r.db.Delete(&models.Employee{}, id)
	return result.Error
}

func (r *employeeRepository) FindUserByID(userID int) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, userID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, result.Error
}
