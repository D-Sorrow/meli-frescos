package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type EmployeeService interface {
	GetEmployees() ([]models.Employee, error)
	GetEmployeeById(employeeId int) (models.Employee, error)
}
