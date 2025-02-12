package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/error_management"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/mappers"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type EmployeeHandler struct {
	service   service.EmployeeService
	validator *validator.Validate
}

func NewEmployeeHandler(service service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		service:   service,
		validator: validator.New(),
	}
}

func (handler *EmployeeHandler) GetEmployees() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employees, err := handler.service.GetEmployees()

		if err != nil {
			error_management.HandleErrorEmployee(w, err)
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
			error_management.HandleErrorEmployee(w, errors.New("ID-DEC-ERR"))
			return
		}

		employee, err := handler.service.GetEmployeeById(id)
		employeeDto := mappers.EmployeeModelToDTO(employee)

		if err != nil {
			error_management.HandleErrorEmployee(w, err)
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
			error_management.HandleErrorEmployee(w, errors.New("BODY-DEC-ERR"))
			return
		}

		if err := handler.validator.Struct(employeeToCreate); err != nil {
			error_management.HandleErrorEmployee(w, err)
			return
		}

		employeeModel := mappers.EmployeeDTOToModel(employeeToCreate)
		employeeCreated, err := handler.service.CreateEmployee(*employeeModel)

		if err != nil {
			error_management.HandleErrorEmployee(w, err)
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
			error_management.HandleErrorEmployee(w, errors.New("ID-DEC-ERR"))
			return
		}

		var employeePatchRequestDTO dto.EmployeePatchRequestDTO
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&employeePatchRequestDTO); err != nil {
			error_management.HandleErrorEmployee(w, errors.New("BODY-DEC-ERR"))
			return
		}

		employeePatchRequestModel := mappers.EmployeePatchRequestDTOToModel(employeePatchRequestDTO)
		employeeUpdated, err := handler.service.UpdateEmployee(id, *employeePatchRequestModel)

		if err != nil {
			error_management.HandleErrorEmployee(w, err)
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
			error_management.HandleErrorEmployee(w, errors.New("ID-DEC-ERR"))
			return
		}

		errorService := handler.service.DeleteEmployee(employeeId)
		if errorService != nil {
			error_management.HandleErrorEmployee(w, err)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Empleado elimindo correctamente",
		})
	}
}
