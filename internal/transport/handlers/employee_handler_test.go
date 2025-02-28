package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	serviceMock "github.com/D-Sorrow/meli-frescos/mocks/internal_/domain/service"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetEmployees(t *testing.T) {
	t.Run("GetEmployees success returning all employees", func(t *testing.T) {
		mockService := new(serviceMock.MockEmployeeService)
		handler := NewEmployeeHandler(mockService)

		employees := []models.Employee{
			{
				Id:           1,
				CardNumberId: "ABCD001",
				FirstName:    "Alejandro",
				LastName:     "Salazar",
				WarehouseId:  1,
			},
			{
				Id:           2,
				CardNumberId: "ABCD002",
				FirstName:    "Monica",
				LastName:     "Arboleda",
				WarehouseId:  1,
			},
		}

		employeesResponseDto := []dto.EmployeeDTO{
			{
				Id:           1,
				CardNumberId: "ABCD001",
				FirstName:    "Alejandro",
				LastName:     "Salazar",
				WarehouseId:  1,
			},
			{
				Id:           2,
				CardNumberId: "ABCD002",
				FirstName:    "Monica",
				LastName:     "Arboleda",
				WarehouseId:  1,
			},
		}

		expectedResponse := dto.EmployeeResponseDto[[]dto.EmployeeDTO]{
			Code: http.StatusOK,
			Msg:  "Success",
			Data: employeesResponseDto,
		}

		mockService.On("GetEmployees").Return(employees, nil)

		req := httptest.NewRequest("GET", "/api/v1/employees", nil)
		rr := httptest.NewRecorder()
		router := chi.NewRouter()
		router.Handle("/api/v1/employees", handler.GetEmployees())
		router.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)
		var response dto.EmployeeResponseDto[[]dto.EmployeeDTO]
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, "Success", response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})

	t.Run("GetEmployees fails when service return error", func(t *testing.T) {
		mockService := new(serviceMock.MockEmployeeService)
		handler := NewEmployeeHandler(mockService)

		expectedResponse := dto.EmployeeResponseDto[[]dto.EmployeeDTO]{
			Code: http.StatusInternalServerError,
			Msg:  "Internal server error",
			Data: nil,
		}

		mockService.On("GetEmployees").Return(nil, service.ErrEmployeeServiceDefault)

		req := httptest.NewRequest("GET", "/api/v1/employees", nil)
		rr := httptest.NewRecorder()
		router := chi.NewRouter()
		router.Handle("/api/v1/employees", handler.GetEmployees())
		router.ServeHTTP(rr, req)

		require.Equal(t, http.StatusInternalServerError, rr.Code)
		var response dto.EmployeeResponseDto[[]dto.EmployeeDTO]
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, "Internal server error", response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})
}

func TestGetEmployeeById(t *testing.T) {
	t.Run("GetEmployeeById success returning the employee", func(t *testing.T) {
		mockService := new(serviceMock.MockEmployeeService)
		handler := NewEmployeeHandler(mockService)

		employee := models.Employee{
			Id:           1,
			CardNumberId: "ABCD001",
			FirstName:    "Alejandro",
			LastName:     "Salazar",
			WarehouseId:  1,
		}

		employeesResponseDto := dto.EmployeeDTO{
			Id:           1,
			CardNumberId: "ABCD001",
			FirstName:    "Alejandro",
			LastName:     "Salazar",
			WarehouseId:  1,
		}

		expectedResponse := dto.EmployeeResponseDto[dto.EmployeeDTO]{
			Code: http.StatusOK,
			Msg:  "Success",
			Data: employeesResponseDto,
		}

		mockService.On("GetEmployeeById", 1).Return(employee, nil)

		req := httptest.NewRequest("GET", "/api/v1/employees/1", nil)
		rr := httptest.NewRecorder()
		router := chi.NewRouter()
		router.Handle("/api/v1/employees/{id}", handler.GetEmployeeById())
		router.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)
		var response dto.EmployeeResponseDto[dto.EmployeeDTO]
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, "Success", response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})

	t.Run("GetEmployeeById fails when id is invalid", func(t *testing.T) {
		mockService := new(serviceMock.MockEmployeeService)
		handler := NewEmployeeHandler(mockService)

		expectedResponse := dto.EmployeeResponseDto[dto.EmployeeDTO]{
			Code: http.StatusBadRequest,
			Msg:  "El formato del ID no es válido",
			Data: dto.EmployeeDTO{},
		}

		req := httptest.NewRequest("GET", "/api/v1/employees/1a", nil)
		rr := httptest.NewRecorder()
		router := chi.NewRouter()
		router.Handle("/api/v1/employees/{id}", handler.GetEmployeeById())
		router.ServeHTTP(rr, req)

		require.Equal(t, http.StatusBadRequest, rr.Code)
		var response dto.EmployeeResponseDto[dto.EmployeeDTO]
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, expectedResponse.Msg, response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})

	t.Run("GetEmployeeById fails when employee is not found", func(t *testing.T) {
		mockService := new(serviceMock.MockEmployeeService)
		handler := NewEmployeeHandler(mockService)

		expectedResponse := dto.EmployeeResponseDto[dto.EmployeeDTO]{
			Code: http.StatusNotFound,
			Msg:  "Empleado no encontrado",
			Data: dto.EmployeeDTO{},
		}

		mockService.On("GetEmployeeById", 1).Return(models.Employee{}, service.ErrEmployeeNotFound)

		req := httptest.NewRequest("GET", "/api/v1/employees/1", nil)
		rr := httptest.NewRecorder()
		router := chi.NewRouter()
		router.Handle("/api/v1/employees/{id}", handler.GetEmployeeById())
		router.ServeHTTP(rr, req)

		require.Equal(t, http.StatusNotFound, rr.Code)
		var response dto.EmployeeResponseDto[dto.EmployeeDTO]
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, expectedResponse.Msg, response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})
}

func TestCreateEmployee(t *testing.T) {
	t.Run("CreateEmployee success returning the employee created", func(t *testing.T) {
		mockService := new(serviceMock.MockEmployeeService)
		handler := NewEmployeeHandler(mockService)

		employeeToCreate := dto.EmployeeRequestDTO{
			CardNumberId: "ABCD001",
			FirstName:    "ALEJANDRO",
			LastName:     "SALAZAR",
			WarehouseId:  1,
		}

		employeeCreated := models.Employee{
			Id:           1,
			CardNumberId: "ABCD001",
			FirstName:    "ALEJANDRO",
			LastName:     "SALAZAR",
			WarehouseId:  1,
		}

		employeesResponseDto := dto.EmployeeDTO{
			Id:           1,
			CardNumberId: "ABCD001",
			FirstName:    "ALEJANDRO",
			LastName:     "SALAZAR",
			WarehouseId:  1,
		}

		expectedResponse := dto.EmployeeResponseDto[dto.EmployeeDTO]{
			Code: http.StatusOK,
			Msg:  "Success",
			Data: employeesResponseDto,
		}

		mockService.On("CreateEmployee", mock.Anything).Return(employeeCreated, nil)

		body, _ := json.Marshal(employeeToCreate)
		req := httptest.NewRequest("POST", "/api/v1/employees", bytes.NewReader(body))
		rr := httptest.NewRecorder()
		router := chi.NewRouter()
		router.Handle("/api/v1/employees", handler.CreateEmployee())
		router.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)
		var response dto.EmployeeResponseDto[dto.EmployeeDTO]
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, "Success", response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})
	t.Run("CreateEmployee fails when the body is malformed", func(t *testing.T) {
		mockService := new(serviceMock.MockEmployeeService)
		handler := NewEmployeeHandler(mockService)

		expectedResponse := dto.EmployeeResponseDto[dto.EmployeeDTO]{
			Code: http.StatusBadRequest,
			Msg:  "El cuerpo de la petición está mal formado",
			Data: dto.EmployeeDTO{},
		}

		body := []byte(`{
    		"card_number_id": "412",
    		"first_names": "Jordan",
    		"last_name": "Hernandez",
   			"warehouse_id": 1
			}   `)
		req := httptest.NewRequest("POST", "/api/v1/employees", bytes.NewReader(body))
		rr := httptest.NewRecorder()
		router := chi.NewRouter()
		router.Handle("/api/v1/employees", handler.CreateEmployee())
		router.ServeHTTP(rr, req)

		require.Equal(t, expectedResponse.Code, rr.Code)
		var response dto.EmployeeResponseDto[dto.EmployeeDTO]
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, expectedResponse.Msg, response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})

	t.Run("CreateEmployee fails when the body is without on of the required fields", func(t *testing.T) {
		mockService := new(serviceMock.MockEmployeeService)
		handler := NewEmployeeHandler(mockService)

		expectedResponse := dto.EmployeeResponseDto[dto.EmployeeDTO]{
			Code: http.StatusUnprocessableEntity,
			Msg:  "Validación fallida:  required FirstName, ",
			Data: dto.EmployeeDTO{},
		}

		body := []byte(`{
    		"card_number_id": "412",
    		"last_name": "Hernandez",
   			"warehouse_id": 1
			}   `)
		req := httptest.NewRequest("POST", "/api/v1/employees", bytes.NewReader(body))
		rr := httptest.NewRecorder()
		router := chi.NewRouter()
		router.Handle("/api/v1/employees", handler.CreateEmployee())
		router.ServeHTTP(rr, req)

		require.Equal(t, expectedResponse.Code, rr.Code)
		var response dto.EmployeeResponseDto[dto.EmployeeDTO]
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, expectedResponse.Msg, response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})

	t.Run("CreateEmployee fails when CardNumberId already exists in another employee", func(t *testing.T) {
		employeeToCreate := dto.EmployeeRequestDTO{
			CardNumberId: "ABCD001",
			FirstName:    "ALEJANDRO",
			LastName:     "SALAZAR",
			WarehouseId:  1,
		}

		mockService := new(serviceMock.MockEmployeeService)
		mockService.On("CreateEmployee", mock.Anything).Return(models.Employee{}, service.ErrEmployeeAlreadyExists)
		handler := NewEmployeeHandler(mockService)

		expectedResponse := dto.EmployeeResponseDto[dto.EmployeeDTO]{
			Code: http.StatusConflict,
			Msg:  "Empleado con ese card ID ya existe",
			Data: dto.EmployeeDTO{},
		}

		body, _ := json.Marshal(employeeToCreate)
		req := httptest.NewRequest("POST", "/api/v1/employees", bytes.NewReader(body))
		rr := httptest.NewRecorder()
		router := chi.NewRouter()
		router.Handle("/api/v1/employees", handler.CreateEmployee())
		router.ServeHTTP(rr, req)

		require.Equal(t, expectedResponse.Code, rr.Code)
		var response dto.EmployeeResponseDto[dto.EmployeeDTO]
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, expectedResponse.Msg, response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateEmployee(t *testing.T) {
	t.Run("UpdateEmployee success returning the employee updated", func(t *testing.T) {
		mockService := new(serviceMock.MockEmployeeService)
		handler := NewEmployeeHandler(mockService)

		cardNumberId := "ABCD002"
		firstName := "ALEJANDRA"
		lastName := "SALAZAR"
		warehouseId := 2

		employeeToUpdate := dto.EmployeePatchRequestDTO{
			CardNumberId: &cardNumberId,
			FirstName:    &firstName,
			LastName:     &lastName,
			WarehouseId:  &warehouseId,
		}

		employeeUpdated := models.Employee{
			Id:           1,
			CardNumberId: "ABCD002",
			FirstName:    "ALEJANDRA",
			LastName:     "SALAZAR",
			WarehouseId:  2,
		}

		employeeResponseDto := dto.EmployeeDTO{
			Id:           1,
			CardNumberId: "ABCD002",
			FirstName:    "ALEJANDRA",
			LastName:     "SALAZAR",
			WarehouseId:  2,
		}

		expectedResponse := dto.EmployeeResponseDto[dto.EmployeeDTO]{
			Code: http.StatusOK,
			Msg:  "Success",
			Data: employeeResponseDto,
		}

		mockService.On("UpdateEmployee", 1, mock.Anything).Return(employeeUpdated, nil)

		body, _ := json.Marshal(employeeToUpdate)
		req := httptest.NewRequest("PATCH", "/api/v1/employees/1", bytes.NewReader(body))
		rr := httptest.NewRecorder()
		router := chi.NewRouter()
		router.Handle("/api/v1/employees/{id}", handler.UpdateEmployee())
		router.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)
		var response dto.EmployeeResponseDto[dto.EmployeeDTO]
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, "Success", response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})

	t.Run("UpdateEmployee fails when employee is not found", func(t *testing.T) {
		mockService := new(serviceMock.MockEmployeeService)
		handler := NewEmployeeHandler(mockService)

		cardNumberId := "ABCD002"
		firstName := "ALEJANDRA"
		lastName := "SALAZAR"
		warehouseId := 2

		employeeToUpdate := dto.EmployeePatchRequestDTO{
			CardNumberId: &cardNumberId,
			FirstName:    &firstName,
			LastName:     &lastName,
			WarehouseId:  &warehouseId,
		}

		expectedResponse := dto.EmployeeResponseDto[dto.EmployeeDTO]{
			Code: http.StatusNotFound,
			Msg:  "Empleado no encontrado",
			Data: dto.EmployeeDTO{},
		}

		mockService.On("UpdateEmployee", 1, mock.Anything).Return(models.Employee{}, service.ErrEmployeeNotFound)

		body, _ := json.Marshal(employeeToUpdate)
		req := httptest.NewRequest("PATCH", "/api/v1/employees/1", bytes.NewReader(body))
		rr := httptest.NewRecorder()
		router := chi.NewRouter()
		router.Handle("/api/v1/employees/{id}", handler.UpdateEmployee())
		router.ServeHTTP(rr, req)

		require.Equal(t, http.StatusNotFound, rr.Code)
		var response dto.EmployeeResponseDto[dto.EmployeeDTO]
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, expectedResponse.Msg, response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})

	t.Run("UpdateEmployee fails when employeeid is malformed", func(t *testing.T) {
		mockService := new(serviceMock.MockEmployeeService)
		handler := NewEmployeeHandler(mockService)

		cardNumberId := "ABCD002"
		firstName := "ALEJANDRA"
		lastName := "SALAZAR"
		warehouseId := 2

		employeeToUpdate := dto.EmployeePatchRequestDTO{
			CardNumberId: &cardNumberId,
			FirstName:    &firstName,
			LastName:     &lastName,
			WarehouseId:  &warehouseId,
		}

		expectedResponse := dto.EmployeeResponseDto[dto.EmployeeDTO]{
			Code: http.StatusBadRequest,
			Msg:  "El formato del ID no es válido",
			Data: dto.EmployeeDTO{},
		}

		body, _ := json.Marshal(employeeToUpdate)
		req := httptest.NewRequest("PATCH", "/api/v1/employees/1a", bytes.NewReader(body))
		rr := httptest.NewRecorder()
		router := chi.NewRouter()
		router.Handle("/api/v1/employees/{id}", handler.UpdateEmployee())
		router.ServeHTTP(rr, req)

		require.Equal(t, expectedResponse.Code, rr.Code)
		var response dto.EmployeeResponseDto[dto.EmployeeDTO]
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, expectedResponse.Msg, response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})

	t.Run("UpdateEmployee fails when body is malformed", func(t *testing.T) {
		mockService := new(serviceMock.MockEmployeeService)
		handler := NewEmployeeHandler(mockService)

		expectedResponse := dto.EmployeeResponseDto[dto.EmployeeDTO]{
			Code: http.StatusBadRequest,
			Msg:  "El cuerpo de la petición está mal formado",
			Data: dto.EmployeeDTO{},
		}

		body := []byte(`{
    		"card_number_ids": "412",
    		"last_name": "Hernandez",
   			"warehouse_id": 1
			}   `)

		req := httptest.NewRequest("PATCH", "/api/v1/employees/1", bytes.NewReader(body))
		rr := httptest.NewRecorder()
		router := chi.NewRouter()
		router.Handle("/api/v1/employees/{id}", handler.UpdateEmployee())
		router.ServeHTTP(rr, req)

		require.Equal(t, expectedResponse.Code, rr.Code)
		var response dto.EmployeeResponseDto[dto.EmployeeDTO]
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, expectedResponse.Msg, response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})
}

func TestDeleteEmployee(t *testing.T) {
	t.Run("DeleteEmployee success when employee exists", func(t *testing.T) {
		mockService := new(serviceMock.MockEmployeeService)
		handler := NewEmployeeHandler(mockService)

		expectedResponse := dto.EmployeeResponseDto[string]{
			Code: http.StatusOK,
			Msg:  "Empleado eliminado correctamente",
			Data: "",
		}

		mockService.On("DeleteEmployee", 1).Return(nil)

		req := httptest.NewRequest("DELETE", "/api/v1/employees/1", nil)
		rr := httptest.NewRecorder()
		router := chi.NewRouter()
		router.Handle("/api/v1/employees/{id}", handler.DeleteEmployee())
		router.ServeHTTP(rr, req)

		require.Equal(t, expectedResponse.Code, rr.Code)
		var response dto.EmployeeResponseDto[string]
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, expectedResponse.Msg, response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})

	t.Run("DeleteEmployee fails when the employee is not found", func(t *testing.T) {
		mockService := new(serviceMock.MockEmployeeService)
		handler := NewEmployeeHandler(mockService)

		expectedResponse := dto.EmployeeResponseDto[string]{
			Code: http.StatusNotFound,
			Msg:  "Empleado no encontrado",
			Data: "",
		}

		mockService.On("DeleteEmployee", 1).Return(service.ErrEmployeeNotFound)

		req := httptest.NewRequest("DELETE", "/api/v1/employees/1", nil)
		rr := httptest.NewRecorder()
		router := chi.NewRouter()
		router.Handle("/api/v1/employees/{id}", handler.DeleteEmployee())
		router.ServeHTTP(rr, req)

		require.Equal(t, expectedResponse.Code, rr.Code)
		var response dto.EmployeeResponseDto[string]
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, expectedResponse.Msg, response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})

	t.Run("DeleteEmployee fails when the id is not valid", func(t *testing.T) {
		mockService := new(serviceMock.MockEmployeeService)
		handler := NewEmployeeHandler(mockService)

		expectedResponse := dto.EmployeeResponseDto[string]{
			Code: http.StatusBadRequest,
			Msg:  "El formato del ID no es válido",
			Data: "",
		}

		req := httptest.NewRequest("DELETE", "/api/v1/employees/1a", nil)
		rr := httptest.NewRecorder()
		router := chi.NewRouter()
		router.Handle("/api/v1/employees/{id}", handler.DeleteEmployee())
		router.ServeHTTP(rr, req)

		require.Equal(t, expectedResponse.Code, rr.Code)
		var response dto.EmployeeResponseDto[string]
		json.NewDecoder(rr.Body).Decode(&response)

		require.Equal(t, expectedResponse.Msg, response.Msg)
		require.Equal(t, expectedResponse, response)
		mockService.AssertExpectations(t)
	})
}
