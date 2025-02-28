package service

import (
	"strconv"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/domain/service/error_management"
)

type EmployeeService struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(repository repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		repository: repository,
	}
}

func (_service *EmployeeService) GetEmployees() (employees []models.Employee, err error) {
	allEmployees, err := _service.repository.GetEmployees()

	if err != nil {
		return nil, error_management.HandleErrorEmployeeService(err)
	}

	for _, employee := range allEmployees {
		employees = append(employees, employee)
	}

	return
}

func (_service *EmployeeService) GetEmployeeById(employeeId int) (employee models.Employee, err error) {
	employee, err = _service.repository.GetEmployeeById(employeeId)

	if err != nil {
		return models.Employee{}, error_management.HandleErrorEmployeeService(err)
	}

	return
}

func (_service *EmployeeService) CreateEmployee(employee models.Employee) (models.Employee, error) {
	allEmployees, err := _service.repository.GetEmployees()
	if err != nil {
		return models.Employee{}, error_management.HandleErrorEmployeeService(err)
	}
	for _, emp := range allEmployees {
		if emp.CardNumberId == employee.CardNumberId {
			return models.Employee{}, service.ErrEmployeeAlreadyExists
		}
	}
	if err = _service.repository.CreateEmployee(&employee); err != nil {
		return models.Employee{}, error_management.HandleErrorEmployeeService(err)
	}
	return employee, nil
}

func (_service *EmployeeService) UpdateEmployee(employeeId int, employee models.EmployeePatchRequest) (employeeUpdated models.Employee, err error) {
	employeeUpdated, err = _service.repository.GetEmployeeById(employeeId)
	allEmployees, errorAll := _service.repository.GetEmployees()
	if errorAll != nil {
		return models.Employee{}, error_management.HandleErrorEmployeeService(err)
	}
	if err != nil {
		return models.Employee{}, error_management.HandleErrorEmployeeService(err)
	}

	if employee.CardNumberId != nil {
		employeeUpdated.CardNumberId = *employee.CardNumberId
		for _, emp := range allEmployees {
			if emp.CardNumberId == *employee.CardNumberId {
				return models.Employee{}, service.ErrEmployeeAlreadyExists
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
	_service.repository.UpdateEmployee(&employeeUpdated)
	return

}

func (_service *EmployeeService) DeleteEmployee(employeeId int) (err error) {
	err = _service.repository.DeleteEmployee(employeeId)

	if err != nil {
		err = error_management.HandleErrorEmployeeService(err)
	}

	return
}

func (_service *EmployeeService) GetReportInboundOrdersByEmployee(employeeId string) (employees []models.EmployeeReportInboundOrders, err error) {
	if employeeId != "" {
		id, err := strconv.Atoi(employeeId)
		if err != nil {
			return nil, error_management.HandleErrorEmployeeService(err)
		}
		employee, err := _service.repository.GetInboundOrdersCountByEmployeeId(id)
		if err != nil {
			return nil, error_management.HandleErrorEmployeeService(err)
		}
		employees = append(employees, employee)
	} else {
		allEmployees, err := _service.repository.GetInboundOrdersCountAllEmployees()

		if err != nil {
			return nil, error_management.HandleErrorEmployeeService(err)
		}

		employees = append(employees, allEmployees...)
	}

	return
}
