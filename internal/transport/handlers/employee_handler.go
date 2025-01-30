package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers/dto"
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
			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"error": "Internal server error",
			})
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
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": "Formato de ID inválido",
			})
			return
		}

		employee, err := handler.service.GetEmployeeById(id)
		employeeDto := mappers.EmployeeModelToDTO(employee)

		if err != nil {
			if err.Error() == "ENF-SV" {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"error": "Empleado no encontrado",
				})
				return
			} else {
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"error": "Internal server error",
				})
				return
			}
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
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": "Bad request - decoding",
			})
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
			if err.Error() == "EAE_SV" {
				response.JSON(w, http.StatusBadRequest, map[string]any{
					"error": "Empleado con ese Card ID ya existe",
				})
				return
			} else {
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"error": "Internal server error",
				})
				return
			}
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
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": "Formato de ID inválido",
			})
			return
		}

		var employeePatchRequestDTO dto.EmployeePatchRequestDTO
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&employeePatchRequestDTO); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": "Bad request - decoding",
			})
			return
		}

		employeePatchRequestModel := mappers.EmployeePatchRequestDTOToModel(employeePatchRequestDTO)
		employeeUpdated, err := handler.service.UpdateEmployee(id, *employeePatchRequestModel)

		if err != nil {
			if err.Error() == "ENF-SV" {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"error": "Empleado no encontrado",
				})
				return
			} else if err.Error() == "EAE-SV" {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"error": "Empleado con ese card id ya existe",
				})
				return
			} else {
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"error": "Internal server error",
				})
				return
			}
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
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": "El formato del id no es válido",
			})
			return
		}

		errorService := handler.service.DeleteEmployee(employeeId)
		if errorService != nil {
			if errorService.Error() == "ENF-SV" {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"error": "El empleado con ese id no existe",
				})
				return
			} else {
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"error": "Internal server error",
				})
				return
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Empleado elimindo correctamente",
		})
	}
}
