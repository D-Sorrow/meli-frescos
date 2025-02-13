package repository

import (
	"database/sql"
	"errors"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
)

type EmployeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (repository *EmployeeRepository) GetEmployees() (employees map[int]models.Employee, err error) {
	employees = make(map[int]models.Employee)
	rows, err := repository.db.Query("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee models.Employee
		err := rows.Scan(&employee.Id, &employee.CardNumberId, &employee.FirstName, &employee.LastName, &employee.WarehouseId)
		if err != nil {
			panic(err)
		}
		employees[employee.Id] = employee
	}
	return
}

func (repository *EmployeeRepository) GetEmployeeById(employeeId int) (employee models.Employee, err error) {
	row := repository.db.QueryRow("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id = ?", employeeId)
	if row.Err() != nil {
		return models.Employee{}, row.Err()
	}

	if err := row.Scan(&employee.Id, &employee.CardNumberId, &employee.FirstName, &employee.LastName, &employee.WarehouseId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Employee{}, errors.New("ENF-DB")
		}
		return models.Employee{}, err
	}
	return
}

func (repository *EmployeeRepository) CreateEmployee(employee models.Employee) models.Employee {
	panic("implement me")
}

func (repository *EmployeeRepository) UpdateEmployee(employeeId int, employee models.Employee) {
	panic("implement me")
}

func (repository *EmployeeRepository) DeleteEmployee(employeeId int) error {
	panic("implement me")
}

func (repository *EmployeeRepository) generateId() int {
	panic("implement me")
}
