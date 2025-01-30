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
			return models.Employee{}, errors.New("ENF-SV")
		} else {
			return models.Employee{}, errors.New("internal server error")
		}
	}

	return
}

func (service *EmployeeService) CreateEmployee(employee models.Employee) (models.Employee, error) {
	allEmployees, err := service.repository.GetEmployees()
	if err != nil {
		return models.Employee{}, errors.New("ERE_DB")
	}
	for _, emp := range allEmployees {
		if emp.CardNumberId == employee.CardNumberId {
			return models.Employee{}, errors.New("EAE_DB")
		}
	}
	return service.repository.CreateEmployee(employee), nil
}

func (service *EmployeeService) UpdateEmployee(employeeId int, employee models.EmployeePatchRequest) (employeeUpdated models.Employee, err error) {
	employeeUpdated, err = service.repository.GetEmployeeById(employeeId)
	allEmployees, errorAll := service.repository.GetEmployees()
	if errorAll != nil {
		return models.Employee{}, errors.New("ISE-SV")
	}
	if err != nil {
		if err.Error() == "ENF-DB" {
			return models.Employee{}, errors.New("ENF-SV")
		} else {
			return models.Employee{}, errors.New("ISE-SV")
		}
	}

	if employee.CardNumberId != nil {
		employeeUpdated.CardNumberId = *employee.CardNumberId
		for _, emp := range allEmployees {
			if emp.CardNumberId == *employee.CardNumberId {
				return models.Employee{}, errors.New("EAE-DB")
			}
		}
	}

	if employee.FirstName != nil {
		employeeUpdated.FirstName = *employee.FirstName
	}

	if employee.LastName != nil {
		employeeUpdated.LastName = *employee.LastName
	}

	if employee.WarehouseId != nil {
		employeeUpdated.WarehouseId = *employee.WarehouseId
	}
	service.repository.UpdateEmployee(employeeId, employeeUpdated)
	return

}
