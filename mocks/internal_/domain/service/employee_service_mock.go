package service_mock

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type MockEmployeeService struct {
	mock.Mock
}

func (m *MockEmployeeService) GetEmployees() ([]models.Employee, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Employee), args.Error(1)
}

func (m *MockEmployeeService) GetEmployeeById(id int) (models.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *MockEmployeeService) CreateEmployee(employee models.Employee) (models.Employee, error) {
	args := m.Called(employee)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *MockEmployeeService) UpdateEmployee(employeeId int, employee models.EmployeePatchRequest) (models.Employee, error) {
	args := m.Called(employeeId, employee)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *MockEmployeeService) DeleteEmployee(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockEmployeeService) GetReportInboundOrdersByEmployee(employeeId string) ([]models.EmployeeReportInboundOrders, error) {
	args := m.Called(employeeId)
	return args.Get(0).([]models.EmployeeReportInboundOrders), args.Error(1)
}
