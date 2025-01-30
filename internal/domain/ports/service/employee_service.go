package service

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type EmployeeService interface {
	GetEmployees() ([]models.Employee, error)
	GetEmployeeById(employeeId int) (models.Employee, error)
	CreateEmployee(employee models.Employee) (models.Employee, error)
	UpdateEmployee(employeeId int, employee models.EmployeePatchRequest) (models.Employee, error)
}
