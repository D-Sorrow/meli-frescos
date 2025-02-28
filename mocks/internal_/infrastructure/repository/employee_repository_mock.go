package repository

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type MockEmployeeRepository struct {
	mock.Mock
}

func (m *MockEmployeeRepository) GetEmployees() (map[int]models.Employee, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[int]models.Employee), args.Error(1)
}

func (m *MockEmployeeRepository) GetEmployeeById(employeeId int) (models.Employee, error) {
	args := m.Called(employeeId)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *MockEmployeeRepository) CreateEmployee(employee *models.Employee) error {
	args := m.Called(employee)
	return args.Error(0)
}

func (m *MockEmployeeRepository) UpdateEmployee(employee *models.Employee) error {
	args := m.Called(employee)
	return args.Error(0)
}

func (m *MockEmployeeRepository) DeleteEmployee(employeeId int) error {
	args := m.Called(employeeId)
	return args.Error(0)
}

func (m *MockEmployeeRepository) GetInboundOrdersCountByEmployeeId(employeeId int) (models.EmployeeReportInboundOrders, error) {
	args := m.Called()
	return args.Get(0).(models.EmployeeReportInboundOrders), args.Error(1)
}

func (m *MockEmployeeRepository) GetInboundOrdersCountAllEmployees() ([]models.EmployeeReportInboundOrders, error) {
	args := m.Called()
	return args.Get(0).([]models.EmployeeReportInboundOrders), args.Error(1)
}
