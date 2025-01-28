package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
)

type EmployeeService struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(repository repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		repository: repository,
	}
}

func (service *EmployeeService) GetEmployees() (employees []models.Employee, err error) {
	//Mapper a DTO de respuesta
	allEmployees, err := service.repository.GetEmployees()

	if err != nil {
		return nil, errors.New("employee-repo-error")
	}

	for _, employee := range allEmployees {
		employees = append(employees, employee)
	}

	return
}

func (service *EmployeeService) GetEmployeeById(employeeId int) (employee models.Employee, err error) {
	employee, err = service.repository.GetEmployeeById(employeeId)

	if err != nil {
		if err.Error() == "ENF-DB" {
			return models.Employee{}, errors.New("employee-not-found")
		} else {
			return models.Employee{}, errors.New("internal server error")
		}
	}

	return
}
