package repository

import "github.com/D-Sorrow/meli-frescos/pkg/models"

type EmployeeRepository interface {
	GetEmployees() (map[int]models.Employee, error)
	GetEmployeeById(employeeId int) (models.Employee, error)
	CreateEmployee(employee models.Employee) error
	UpdateEmployee(employee models.Employee) (models.Employee, error)
	DeleteEmployee(employeeId int) error
}
