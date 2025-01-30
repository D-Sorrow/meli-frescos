package repository

import "github.com/D-Sorrow/meli-frescos/internal/domain/models"

type EmployeeRepository interface {
	GetEmployees() (map[int]models.Employee, error)
	GetEmployeeById(employeeId int) (models.Employee, error)
	CreateEmployee(employee models.Employee) models.Employee
	UpdateEmployee(employeeId int, employee models.Employee)
}
