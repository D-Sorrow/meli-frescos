package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

type EmployeeHandler struct {
	service service.EmployeeService
}

func NewEmployeeHandler(service service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (handler *EmployeeHandler) GetEmployees() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employees, err := handler.service.GetEmployees()

		if err != nil {
			error_management.HandleErrorEmployee(w, err.Error(), nil)
			return
		}
		var employeesDto []dto.EmployeeDTO
		for _, employee := range employees {
			employeeDto := mappers.EmployeeModelToDTO(employee)
			employeesDto = append(employeesDto, *employeeDto)
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": employeesDto,
		})
	}
}

func (handler *EmployeeHandler) GetEmployeeById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			error_management.HandleErrorEmployee(w, "ID-DEC-ERR", nil)
			return
		}

		employee, err := handler.service.GetEmployeeById(id)
		employeeDto := mappers.EmployeeModelToDTO(employee)

		if err != nil {
			error_management.HandleErrorEmployee(w, err.Error(), nil)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": employeeDto,
		})

	}
}

func (handler *EmployeeHandler) CreateEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var employeeToCreate dto.EmployeeRequestDTO
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&employeeToCreate); err != nil {
			error_management.HandleErrorEmployee(w, "BODY-ERR", err)
			return
		}

		if employeeToCreate.CardNumberId == 0 || employeeToCreate.FirstName == "" || employeeToCreate.LastName == "" || employeeToCreate.WarehouseId == 0 {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": "Bad request - Campos incompletos",
			})
			return
		}

		employeeModel := mappers.EmployeeDTOToModel(employeeToCreate)
		employeeCreated, err := handler.service.CreateEmployee(*employeeModel)

		if err != nil {
			error_management.HandleErrorEmployee(w, err.Error(), nil)
			return
		}

		employeeDto := mappers.EmployeeModelToDTO(employeeCreated)
		response.JSON(w, http.StatusCreated, map[string]any{
			"data": employeeDto,
		})
	}
}

func (handler *EmployeeHandler) UpdateEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			error_management.HandleErrorEmployee(w, "ID-DEC-ERR", nil)
			return
		}

		var employeePatchRequestDTO dto.EmployeePatchRequestDTO
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&employeePatchRequestDTO); err != nil {
			error_management.HandleErrorEmployee(w, "BODY-ERR", err)
			return
		}

		employeePatchRequestModel := mappers.EmployeePatchRequestDTOToModel(employeePatchRequestDTO)
		employeeUpdated, err := handler.service.UpdateEmployee(id, *employeePatchRequestModel)

		if err != nil {
			error_management.HandleErrorEmployee(w, err.Error(), nil)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": employeeUpdated,
		})
	}
}

func (handler *EmployeeHandler) DeleteEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employeeIdString := chi.URLParam(r, "id")
		employeeId, err := strconv.Atoi(employeeIdString)
		if err != nil {
			error_management.HandleErrorEmployee(w, "ID-DEC-ERR", nil)
			return
		}

		errorService := handler.service.DeleteEmployee(employeeId)
		if errorService != nil {
			error_management.HandleErrorEmployee(w, errorService.Error(), nil)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Empleado elimindo correctamente",
		})
	}
}
