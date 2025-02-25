package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type EmployeeService interface {
	GetEmployees() ([]models.Employee, error)
	GetEmployeeById(employeeId int) (models.Employee, error)
	CreateEmployee(employee models.Employee) (models.Employee, error)
	UpdateEmployee(employeeId int, employee models.EmployeePatchRequest) (models.Employee, error)
	DeleteEmployee(employeeId int) error
	GetReportInboundOrdersByEmployee(employeeId string) ([]models.EmployeeReportInboundOrders, error)
}

var (
	ErrEmployeeNotFound       = errors.New("employee with the id provided not found in the database")
	ErrEmployeeDecodingError  = errors.New("error decoding id")
	ErrEmployeeServiceDefault = errors.New("internal server error")
	ErrEmployeeAlreadyExists  = errors.New("employee already exists in the database")
)
