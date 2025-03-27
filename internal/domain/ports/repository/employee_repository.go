package repository

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type EmployeeRepository interface {
	GetEmployees() (map[int]models.Employee, error)
	GetEmployeeById(employeeId int) (models.Employee, error)
	CreateEmployee(employee *models.Employee) error
	UpdateEmployee(employee *models.Employee) error
	DeleteEmployee(employeeId int) error
	GetInboundOrdersCountByEmployeeId(employeeId int) (models.EmployeeReportInboundOrders, error)
	GetInboundOrdersCountAllEmployees() ([]models.EmployeeReportInboundOrders, error)
}

var (
	ErrEmployeeNotFound            = errors.New("employee not found in the database with the id provided")
	ErrEmployeeInternalServerError = errors.New("repository internal server error")
)
