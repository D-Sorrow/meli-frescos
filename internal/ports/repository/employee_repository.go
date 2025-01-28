package repository

import "github.com/D-Sorrow/meli-frescos/pkg/models"

type EmployeeRepository interface {
	GetEmployees() (map[int]models.Employee, error)
	GetEmployeeById(employeeId int) (models.Employee, error)
}
