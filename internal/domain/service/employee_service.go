package service

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	"github.com/D-Sorrow/meli-frescos/pkg/mappers"
)

type EmployeeService struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(repository repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		repository: repository,
	}
}

func (service *EmployeeService) GetEmployees() (employeesDto []dto.EmployeeDTO, err error) {
	//Mapper a DTO de respuesta
	employees, err := service.repository.GetEmployees()

	if err != nil {
		return nil, errors.New("employee-repo-error")
	}

	for _, employee := range employees {
		employeeDto := mappers.EmployeeModelToDTO(employee)
		employeesDto = append(employeesDto, *employeeDto)
	}

	return
}

func (service *EmployeeService) GetEmployeeById(employeeId int) (employee dto.EmployeeDTO, err error) {
	employeeFound, err := service.repository.GetEmployeeById(employeeId)

	if err != nil {
		if err.Error() == "ENF-DB" {
			return dto.EmployeeDTO{}, errors.New("employee-not-found")
		} else {
			return dto.EmployeeDTO{}, errors.New("internal server error")
		}
	}

	employee = *mappers.EmployeeModelToDTO(employeeFound)

	return
}
