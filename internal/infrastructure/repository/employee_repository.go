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
