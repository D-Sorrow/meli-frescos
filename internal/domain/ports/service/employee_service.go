package service

import "github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"

type EmployeeService interface {
	GetEmployees() ([]dto.EmployeeDTO, error)
	GetEmployeeById(employeeId int) (dto.EmployeeDTO, error)
}
