package repository

import (
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type EmployeeRepository struct {
	db map[int]models.Employee
}

func NewEmployeeRepository(db map[int]models.Employee) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (repository *EmployeeRepository) GetEmployees() (employees map[int]models.Employee, err error) {
	employees = make(map[int]models.Employee)
	for key, value := range repository.db {
		employees[key] = value
	}

	return
}

func (repository *EmployeeRepository) GetEmployeeById(employeeId int) (employee models.Employee, err error) {
	employee, exists := repository.db[employeeId]
	if !exists {
		return models.Employee{}, errors.New("ENF-DB")
	}

	return employee, nil
}

func (repository *EmployeeRepository) CreateEmployee(employee models.Employee) models.Employee {
	id := repository.generateId()
	employeeCreated := models.Employee{
		Id:           id,
		CardNumberId: employee.CardNumberId,
		FirstName:    employee.FirstName,
		LastName:     employee.LastName,
		WarehouseId:  employee.WarehouseId,
	}
	repository.db[id] = employeeCreated
	return employeeCreated
}

func (repository *EmployeeRepository) UpdateEmployee(employeeId int, employee models.Employee) {
	repository.db[employeeId] = employee
}

func (repository *EmployeeRepository) DeleteEmployee(employeeId int) error {
	_, exists := repository.db[employeeId]
	if !exists {
		return errors.New("ENF-DB")
	}

	delete(repository.db, employeeId)
	return nil
}

func (repository *EmployeeRepository) generateId() int {
	maxId := 0
	for id := range repository.db {
		if id > maxId {
			maxId = id
		}
	}

	return int(maxId + 1)
}
