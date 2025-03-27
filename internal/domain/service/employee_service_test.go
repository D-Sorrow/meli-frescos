package service

import (
	"testing"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/repository"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	mockRepo "github.com/D-Sorrow/meli-frescos/mocks/internal_/infrastructure/repository"
	"github.com/stretchr/testify/require"
)

func TestGetEmployees(t *testing.T) {
	t.Run("GetAllEmployees success returns all employees", func(t *testing.T) {
		mockRepo := new(mockRepo.MockEmployeeRepository)

		employeesRepoResponse := map[int]models.Employee{
			1: {
				Id:           1,
				CardNumberId: "ABCD001",
				FirstName:    "ALEJANDRO",
				LastName:     "SALAZAR",
				WarehouseId:  1,
			},
			2: {
				Id:           2,
				CardNumberId: "ABCD002",
				FirstName:    "ALEJANDRA",
				LastName:     "GARCIA",
				WarehouseId:  2,
			},
		}

		employeesServiceResponseExpected := []models.Employee{
			{
				Id:           1,
				CardNumberId: "ABCD001",
				FirstName:    "ALEJANDRO",
				LastName:     "SALAZAR",
				WarehouseId:  1,
			},
			{
				Id:           2,
				CardNumberId: "ABCD002",
				FirstName:    "ALEJANDRA",
				LastName:     "GARCIA",
				WarehouseId:  2,
			},
		}

		mockRepo.On("GetEmployees").Return(employeesRepoResponse, nil)

		serv := NewEmployeeService(mockRepo)

		employeesActual, err := serv.GetEmployees()

		require.NoError(t, err)
		require.Equal(t, employeesServiceResponseExpected, employeesActual)
	})

	t.Run("GetAllEmployees fails due to internal server error", func(t *testing.T) {
		mockRepo := new(mockRepo.MockEmployeeRepository)

		mockRepo.On("GetEmployees").Return(nil, repository.ErrEmployeeInternalServerError)

		serv := NewEmployeeService(mockRepo)

		employeesActual, err := serv.GetEmployees()

		require.Error(t, err)
		require.ErrorIs(t, err, service.ErrEmployeeServiceDefault)
		require.Nil(t, employeesActual)
	})
}

func TestGetEmployeeById(t *testing.T) {
	t.Run("GetEmployeeById success return the employee", func(t *testing.T) {
		mockRepo := new(mockRepo.MockEmployeeRepository)

		employeeRepoResponse := models.Employee{
			Id:           1,
			CardNumberId: "ABCD001",
			FirstName:    "ALEJANDRO",
			LastName:     "SALAZAR",
			WarehouseId:  1,
		}

		employeeServiceResponseExpected := models.Employee{
			Id:           1,
			CardNumberId: "ABCD001",
			FirstName:    "ALEJANDRO",
			LastName:     "SALAZAR",
			WarehouseId:  1,
		}

		mockRepo.On("GetEmployeeById", 1).Return(employeeRepoResponse, nil)

		serv := NewEmployeeService(mockRepo)

		employeesActual, err := serv.GetEmployeeById(1)

		require.NoError(t, err)
		require.Equal(t, employeeServiceResponseExpected, employeesActual)
	})

	t.Run("GetEmployeeById fails returning a NotFoundError", func(t *testing.T) {
		mockRepo := new(mockRepo.MockEmployeeRepository)

		mockRepo.On("GetEmployeeById", 2).Return(models.Employee{}, repository.ErrEmployeeNotFound)

		serv := NewEmployeeService(mockRepo)

		employeesActual, err := serv.GetEmployeeById(2)

		require.Error(t, err)
		require.ErrorIs(t, err, service.ErrEmployeeNotFound)
		require.Empty(t, employeesActual)
	})
}

func TestCreateEmployee(t *testing.T) {
	t.Run("CreateEmployee success returning the employee created", func(t *testing.T) {
		newEmployee := models.Employee{
			Id:           1,
			CardNumberId: "ABCD001",
			FirstName:    "ALEJANDRO",
			LastName:     "SALAZAR",
			WarehouseId:  1,
		}
		mockRepo := new(mockRepo.MockEmployeeRepository)
		mockRepo.On("GetEmployees").Return(nil, nil)
		mockRepo.On("CreateEmployee", &newEmployee).Return(nil)

		serv := NewEmployeeService(mockRepo)

		employeeCreated, err := serv.CreateEmployee(newEmployee)

		require.NoError(t, err)
		require.Equal(t, newEmployee, employeeCreated)
	})

	t.Run("CreateEmployee fails when GetEmployees return error", func(t *testing.T) {
		newEmployee := models.Employee{
			Id:           1,
			CardNumberId: "ABCD001",
			FirstName:    "ALEJANDRO",
			LastName:     "SALAZAR",
			WarehouseId:  1,
		}
		mockRepo := new(mockRepo.MockEmployeeRepository)
		mockRepo.On("GetEmployees").Return(nil, repository.ErrEmployeeInternalServerError)
		mockRepo.On("CreateEmployee", &newEmployee).Return(nil)

		serv := NewEmployeeService(mockRepo)

		employeeCreated, err := serv.CreateEmployee(newEmployee)

		require.Error(t, err)
		require.ErrorIs(t, err, service.ErrEmployeeServiceDefault)
		require.Empty(t, employeeCreated)
	})

	t.Run("CreateEmployee fails when new employee CardNumberId already exists", func(t *testing.T) {
		newEmployee := models.Employee{
			Id:           2,
			CardNumberId: "ABCD001",
			FirstName:    "ALEJANDRA",
			LastName:     "GARCIA",
			WarehouseId:  1,
		}

		existingEmployees := map[int]models.Employee{
			1: {
				Id:           1,
				CardNumberId: "ABCD001",
				FirstName:    "ALEJANDRO",
				LastName:     "SALAZAR",
				WarehouseId:  1,
			},
		}
		mockRepo := new(mockRepo.MockEmployeeRepository)
		mockRepo.On("GetEmployees").Return(existingEmployees, nil)
		mockRepo.On("CreateEmployee", &newEmployee).Return(nil)

		serv := NewEmployeeService(mockRepo)

		employeeCreated, err := serv.CreateEmployee(newEmployee)

		require.Error(t, err)
		require.ErrorIs(t, err, service.ErrEmployeeAlreadyExists)
		require.Empty(t, employeeCreated)
	})

	t.Run("CreateEmployee fails when repository returns an error", func(t *testing.T) {
		newEmployee := models.Employee{
			Id:           1,
			CardNumberId: "ABCD001",
			FirstName:    "ALEJANDRO",
			LastName:     "SALAZAR",
			WarehouseId:  1,
		}
		mockRepo := new(mockRepo.MockEmployeeRepository)
		mockRepo.On("GetEmployees").Return(nil, nil)
		mockRepo.On("CreateEmployee", &newEmployee).Return(repository.ErrEmployeeInternalServerError)

		serv := NewEmployeeService(mockRepo)

		employeeCreated, err := serv.CreateEmployee(newEmployee)

		require.Error(t, err)
		require.ErrorIs(t, err, service.ErrEmployeeServiceDefault)
		require.Empty(t, employeeCreated)
	})

}

func TestUpdateEmployee(t *testing.T) {
	t.Run("UpdateEmployee success with all the parameters returning the employee updated", func(t *testing.T) {

		cardNumberId := "ABCD001"
		firstName := "JUANA"
		lastName := "DE ARCO"
		warehouseId := 2

		employeePatchToUpdate := models.EmployeePatchRequest{
			CardNumberId: &cardNumberId,
			FirstName:    &firstName,
			LastName:     &lastName,
			WarehouseId:  &warehouseId,
		}

		employeeToUpdate := models.Employee{
			Id:           1,
			CardNumberId: "ABCD001",
			FirstName:    "JUANA",
			LastName:     "DE ARCO",
			WarehouseId:  2,
		}

		existingEmployees := map[int]models.Employee{
			1: {
				Id:           1,
				CardNumberId: "ABCD001",
				FirstName:    "ALEJANDRO",
				LastName:     "SALAZAR",
				WarehouseId:  1,
			},
		}

		expectedUpdatedEmployee := models.Employee{
			Id:           1,
			CardNumberId: "ABCD001",
			FirstName:    "JUANA",
			LastName:     "DE ARCO",
			WarehouseId:  2,
		}

		mockRepo := new(mockRepo.MockEmployeeRepository)
		mockRepo.On("GetEmployeeById", 1).Return(existingEmployees[1], nil)
		mockRepo.On("GetEmployees").Return(existingEmployees, nil)
		mockRepo.On("UpdateEmployee", &employeeToUpdate).Return(nil)

		serv := NewEmployeeService(mockRepo)

		employeeUpdated, err := serv.UpdateEmployee(1, employeePatchToUpdate)

		require.NoError(t, err)
		require.Equal(t, expectedUpdatedEmployee, employeeUpdated)
	})

	t.Run("UpdateEmployee fails when employee not exists", func(t *testing.T) {
		cardNumberId := "ABCD001"
		firstName := "JUANA"
		lastName := "DE ARCO"
		warehouseId := 2

		employeePatchToUpdate := models.EmployeePatchRequest{
			CardNumberId: &cardNumberId,
			FirstName:    &firstName,
			LastName:     &lastName,
			WarehouseId:  &warehouseId,
		}
		existingEmployees := map[int]models.Employee{
			1: {
				Id:           1,
				CardNumberId: "ABCD001",
				FirstName:    "ALEJANDRO",
				LastName:     "SALAZAR",
				WarehouseId:  1,
			},
		}

		mockRepo := new(mockRepo.MockEmployeeRepository)
		mockRepo.On("GetEmployeeById", 2).Return(models.Employee{}, repository.ErrEmployeeNotFound)
		mockRepo.On("GetEmployees").Return(existingEmployees, nil)

		serv := NewEmployeeService(mockRepo)

		employeeUpdated, err := serv.UpdateEmployee(2, employeePatchToUpdate)

		require.Error(t, err)
		require.ErrorIs(t, err, service.ErrEmployeeNotFound)
		require.Empty(t, employeeUpdated)

		mockRepo.AssertNotCalled(t, "UpdateEmployee", 2)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateEmployee fails when GetEmployees return error", func(t *testing.T) {
		cardNumberId := "ABCD001"
		firstName := "JUANA"
		lastName := "DE ARCO"
		warehouseId := 2

		employeePatchToUpdate := models.EmployeePatchRequest{
			CardNumberId: &cardNumberId,
			FirstName:    &firstName,
			LastName:     &lastName,
			WarehouseId:  &warehouseId,
		}

		mockRepo := new(mockRepo.MockEmployeeRepository)
		mockRepo.On("GetEmployees").Return(nil, repository.ErrEmployeeInternalServerError)

		serv := NewEmployeeService(mockRepo)

		employeeUpdated, err := serv.UpdateEmployee(2, employeePatchToUpdate)

		require.Error(t, err)
		require.ErrorIs(t, err, service.ErrEmployeeServiceDefault)
		require.Empty(t, employeeUpdated)

		mockRepo.AssertNotCalled(t, "UpdateEmployee")
		mockRepo.AssertNotCalled(t, "GetEmployeeById")
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateEmployee fails when the CarNumberId already exists in other employee", func(t *testing.T) {

		cardNumberId := "ABCD002"
		firstName := "JUANA"
		lastName := "DE ARCO"
		warehouseId := 2

		employeePatchToUpdate := models.EmployeePatchRequest{
			CardNumberId: &cardNumberId,
			FirstName:    &firstName,
			LastName:     &lastName,
			WarehouseId:  &warehouseId,
		}

		existingEmployees := map[int]models.Employee{
			1: {
				Id:           1,
				CardNumberId: "ABCD001",
				FirstName:    "ALEJANDRO",
				LastName:     "SALAZAR",
				WarehouseId:  1,
			},
			2: {
				Id:           2,
				CardNumberId: "ABCD002",
				FirstName:    "ALEJANDRA",
				LastName:     "GARCIA",
				WarehouseId:  2,
			},
		}

		mockRepo := new(mockRepo.MockEmployeeRepository)
		mockRepo.On("GetEmployeeById", 1).Return(existingEmployees[1], nil)
		mockRepo.On("GetEmployees").Return(existingEmployees, nil)

		serv := NewEmployeeService(mockRepo)

		employeeUpdated, err := serv.UpdateEmployee(1, employeePatchToUpdate)

		require.Error(t, err)
		require.ErrorIs(t, err, service.ErrEmployeeAlreadyExists)
		require.Empty(t, employeeUpdated)

		mockRepo.AssertNotCalled(t, "UpdateEmployee")
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteEmployee(t *testing.T) {
	t.Run("DeleteEmployee success returning no error", func(t *testing.T) {
		mockRepo := new(mockRepo.MockEmployeeRepository)
		mockRepo.On("DeleteEmployee", 1).Return(nil)

		serv := NewEmployeeService(mockRepo)
		err := serv.DeleteEmployee(1)

		require.NoError(t, err)
	})

	t.Run("DeleteEmployee fails when the employee doesnt exists", func(t *testing.T) {
		mockRepo := new(mockRepo.MockEmployeeRepository)
		mockRepo.On("DeleteEmployee", 1).Return(repository.ErrEmployeeNotFound)

		serv := NewEmployeeService(mockRepo)
		err := serv.DeleteEmployee(1)

		require.Error(t, err)
		require.ErrorIs(t, err, service.ErrEmployeeNotFound)
	})
}

func TestGetReportInboundOrdersByEmployee(t *testing.T) {
	t.Run("GetReportInboundOrdersByEmployee success with valid employeeId", func(t *testing.T) {
		mockRepo := new(mockRepo.MockEmployeeRepository)

		employeeReport := models.EmployeeReportInboundOrders{
			Id:                1,
			CardNumberId:      "abcd1",
			FirstName:         "Alejo",
			LastName:          "salazar",
			WarehouseId:       1,
			InboundOrderCount: 1,
		}

		mockRepo.On("GetInboundOrdersCountByEmployeeId", 1).Return(employeeReport, nil)

		serv := NewEmployeeService(mockRepo)

		employeesActual, err := serv.GetReportInboundOrdersByEmployee("1")

		require.NoError(t, err)
		require.Equal(t, []models.EmployeeReportInboundOrders{employeeReport}, employeesActual)
	})

	t.Run("GetReportInboundOrdersByEmployee success with empty employeeId", func(t *testing.T) {
		mockRepo := new(mockRepo.MockEmployeeRepository)

		allEmployeesReport := []models.EmployeeReportInboundOrders{
			{
				Id:                1,
				CardNumberId:      "abcd1",
				FirstName:         "Alejo",
				LastName:          "salazar",
				WarehouseId:       1,
				InboundOrderCount: 1,
			},
			{
				Id:                2,
				CardNumberId:      "abcd2",
				FirstName:         "Aleja",
				LastName:          "garcia",
				WarehouseId:       2,
				InboundOrderCount: 2,
			},
		}

		mockRepo.On("GetInboundOrdersCountAllEmployees").Return(allEmployeesReport, nil)

		serv := NewEmployeeService(mockRepo)

		employeesActual, err := serv.GetReportInboundOrdersByEmployee("")

		require.NoError(t, err)
		require.Equal(t, allEmployeesReport, employeesActual)
	})

	t.Run("GetReportInboundOrdersByEmployee fails with invalid employeeId", func(t *testing.T) {
		mockRepo := new(mockRepo.MockEmployeeRepository)

		serv := NewEmployeeService(mockRepo)

		employeesActual, err := serv.GetReportInboundOrdersByEmployee("invalid")

		require.Error(t, err)
		require.ErrorIs(t, err, service.ErrEmployeeDecodingError)
		require.Nil(t, employeesActual)
	})

	t.Run("GetReportInboundOrdersByEmployee fails when repository returns an error for specific employee", func(t *testing.T) {
		mockRepo := new(mockRepo.MockEmployeeRepository)

		mockRepo.On("GetInboundOrdersCountByEmployeeId", 1).Return(models.EmployeeReportInboundOrders{}, repository.ErrEmployeeInternalServerError)

		serv := NewEmployeeService(mockRepo)

		employeesActual, err := serv.GetReportInboundOrdersByEmployee("1")

		require.Error(t, err)
		require.ErrorIs(t, err, service.ErrEmployeeServiceDefault)
		require.Nil(t, employeesActual)
	})

	t.Run("GetReportInboundOrdersByEmployee fails when repository returns an error for all employees", func(t *testing.T) {
		mockRepo := new(mockRepo.MockEmployeeRepository)

		mockRepo.On("GetInboundOrdersCountAllEmployees").Return([]models.EmployeeReportInboundOrders{}, repository.ErrEmployeeInternalServerError)

		serv := NewEmployeeService(mockRepo)

		employeesActual, err := serv.GetReportInboundOrdersByEmployee("")

		require.Error(t, err)
		require.ErrorIs(t, err, service.ErrEmployeeServiceDefault)
		require.Nil(t, employeesActual)
	})
}
