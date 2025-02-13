package repository

import (
	"database/sql"
	"errors"
	"fmt"

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

func (repository *EmployeeRepository) CreateEmployee(employee *models.Employee) error {
	employeeId, err := repository.generateId()
	if err != nil {
		return err
	}
	_, err = repository.db.Exec(
		"INSERT INTO employees (`id`, `card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES (?, ?, ?, ?, ?)",
		employeeId, (*employee).CardNumberId, (*employee).FirstName, (*employee).LastName, (*employee).WarehouseId,
	)
	if err != nil {
		return err
	}
	(*employee).Id = employeeId
	return nil
}

func (repository *EmployeeRepository) UpdateEmployee(employee *models.Employee) error {
	_, err := repository.db.Exec(
		"UPDATE employees SET `card_number_id` = ?, `first_name` = ?, `last_name` = ?, `warehouse_id` = ? WHERE `id` = ?",
		(*employee).CardNumberId, (*employee).FirstName, (*employee).LastName, (*employee).WarehouseId, (*employee).Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repository *EmployeeRepository) DeleteEmployee(employeeId int) error {
	_, err := repository.db.Exec(
		"DELETE FROM employees WHERE id=?",
		employeeId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repository *EmployeeRepository) generateId() (int, error) {
	var maxID int
	err := repository.db.QueryRow("SELECT COALESCE(MAX(id), 0) FROM employees").Scan(&maxID)
	if err != nil {
		return 0, fmt.Errorf("failed to get max ID from table: %v", err)
	}
	newID := maxID + 1
	return newID, nil
}
