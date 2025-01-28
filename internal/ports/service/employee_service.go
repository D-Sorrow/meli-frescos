package service

import (
	"github.com/D-Sorrow/meli-frescos/pkg/dto"
)

type EmployeeService interface {
	GetEmployees() ([]dto.EmployeeDTO, error)
	GetEmployeeById(employeeId int) (dto.EmployeeDTO, error)
}
