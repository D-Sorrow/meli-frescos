package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository/error_management"
)

type EmployeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (_repository *EmployeeRepository) GetEmployees() (employees map[int]models.Employee, err error) {
	employees = make(map[int]models.Employee)
	rows, err := _repository.db.Query("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees")
	if err != nil {
		return nil, error_management.HandleRepositoryError(repository.ErrEmployeeInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee models.Employee
		err := rows.Scan(&employee.Id, &employee.CardNumberId, &employee.FirstName, &employee.LastName, &employee.WarehouseId)
		if err != nil {
			return nil, error_management.HandleRepositoryError(repository.ErrEmployeeInternalServerError, err)
		}
		employees[employee.Id] = employee
	}
	return
}

func (_repository *EmployeeRepository) GetEmployeeById(employeeId int) (employee models.Employee, err error) {
	row := _repository.db.QueryRow("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id = ?", employeeId)
	if row.Err() != nil {
		return models.Employee{}, error_management.HandleRepositoryError(repository.ErrEmployeeInternalServerError, err)
	}

	if err := row.Scan(&employee.Id, &employee.CardNumberId, &employee.FirstName, &employee.LastName, &employee.WarehouseId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Employee{}, error_management.HandleRepositoryError(repository.ErrEmployeeNotFound, err)
		}
		return models.Employee{}, err
	}
	return
}

func (_repository *EmployeeRepository) CreateEmployee(employee *models.Employee) error {
	employeeId, err := _repository.generateId()
	if err != nil {
		return error_management.HandleRepositoryError(repository.ErrEmployeeInternalServerError, err)

	}
	_, err = _repository.db.Exec(
		"INSERT INTO employees (`id`, `card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES (?, ?, ?, ?, ?)",
		employeeId, (*employee).CardNumberId, (*employee).FirstName, (*employee).LastName, (*employee).WarehouseId,
	)
	if err != nil {
		return error_management.HandleRepositoryError(repository.ErrEmployeeInternalServerError, err)
	}
	(*employee).Id = employeeId
	return nil
}

func (_repository *EmployeeRepository) UpdateEmployee(employee *models.Employee) error {
	_, err := _repository.db.Exec(
		"UPDATE employees SET `card_number_id` = ?, `first_name` = ?, `last_name` = ?, `warehouse_id` = ? WHERE `id` = ?",
		(*employee).CardNumberId, (*employee).FirstName, (*employee).LastName, (*employee).WarehouseId, (*employee).Id,
	)
	if err != nil {
		return error_management.HandleRepositoryError(repository.ErrEmployeeInternalServerError, err)
	}
	return nil
}

func (_repository *EmployeeRepository) DeleteEmployee(employeeId int) error {
	result, err := _repository.db.Exec(
		"DELETE FROM employees WHERE id=?",
		employeeId,
	)
	if err != nil {
		return error_management.HandleRepositoryError(repository.ErrEmployeeInternalServerError, err)
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return error_management.HandleRepositoryError(repository.ErrEmployeeNotFound, errors.New("no affected rows"))
	}
	if err != nil {
		return error_management.HandleRepositoryError(repository.ErrEmployeeInternalServerError, err)
	}

	return nil
}

func (_repository *EmployeeRepository) generateId() (int, error) {
	var maxID int

	err := _repository.db.QueryRow("SELECT COALESCE(MAX(id), 0) FROM employees").Scan(&maxID)
	if err != nil {
		return 0, error_management.HandleRepositoryError(repository.ErrEmployeeInternalServerError, err)
	}
	newID := maxID + 1
	return newID, nil
}

func (_repository *EmployeeRepository) GetInboundOrdersCountByEmployeeId(employeeId int) (employee models.EmployeeReportInboundOrders, err error) {
	row := _repository.db.QueryRow("SELECT employees.id, employees.card_number_id, employees.first_name, employees.last_name, employees.warehouse_id, COUNT(inbound_orders.id) AS inbound_orders_count FROM employees LEFT JOIN inbound_orders ON employees.id = inbound_orders.employe_id WHERE employees.id = ? GROUP BY employees.id", employeeId)
	if row.Err() != nil {
		return models.EmployeeReportInboundOrders{}, error_management.HandleRepositoryError(repository.ErrEmployeeInternalServerError, err)

	}

	if err := row.Scan(&employee.Id, &employee.CardNumberId, &employee.FirstName, &employee.LastName, &employee.WarehouseId, &employee.InboundOrderCount); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.EmployeeReportInboundOrders{}, error_management.HandleRepositoryError(repository.ErrEmployeeNotFound, err)

		}
		log.Printf("%v - error: %v", repository.ErrEmployeeInternalServerError.Error(), err.Error())
		return models.EmployeeReportInboundOrders{}, error_management.HandleRepositoryError(repository.ErrEmployeeInternalServerError, err)
	}
	return
}

func (_repository *EmployeeRepository) GetInboundOrdersCountAllEmployees() (employees []models.EmployeeReportInboundOrders, err error) {
	rows, err := _repository.db.Query("SELECT employees.id, employees.card_number_id, employees.first_name, employees.last_name, employees.warehouse_id, COUNT(inbound_orders.id) AS inbound_orders_count FROM employees LEFT JOIN inbound_orders ON employees.id = inbound_orders.employe_id GROUP BY employees.id")
	if err != nil {
		return nil, error_management.HandleRepositoryError(repository.ErrEmployeeInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee models.EmployeeReportInboundOrders
		err = rows.Scan(&employee.Id, &employee.CardNumberId, &employee.FirstName, &employee.LastName, &employee.WarehouseId, &employee.InboundOrderCount)
		if err != nil {
			return nil, error_management.HandleRepositoryError(repository.ErrEmployeeInternalServerError, err)
		}
		employees = append(employees, employee)
	}
	return employees, nil
}
