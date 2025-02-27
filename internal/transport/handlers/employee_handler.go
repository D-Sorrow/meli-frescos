package handlers

import (
	"encoding/json"
	"fmt"
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
			handler.handleError(w, err)
			return
		}
		var employeesDto []dto.EmployeeDTO
		for _, employee := range employees {
			employeeDto := mappers.EmployeeModelToDTO(employee)
			employeesDto = append(employeesDto, *employeeDto)
		}

		respondWithJSON(w, http.StatusOK, "Success", employeesDto)
	}
}

func (handler *EmployeeHandler) GetEmployeeById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			handler.handleError(w, err)
			return
		}

		employee, err := handler.service.GetEmployeeById(id)
		employeeDto := mappers.EmployeeModelToDTO(employee)

		if err != nil {
			handler.handleError(w, err)
			return
		}

		respondWithJSON(w, http.StatusOK, "Success", employeeDto)
	}
}

func (handler *EmployeeHandler) CreateEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var employeeToCreate dto.EmployeeRequestDTO
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&employeeToCreate); err != nil {
			handler.handleError(w, error_management.ErrEmployeeBodyDecoding)
			return
		}

		if err := handler.validator.Struct(employeeToCreate); err != nil {
			handler.handleError(w, err)
			return
		}

		employeeModel := mappers.EmployeeDTOToModel(employeeToCreate)
		employeeCreated, err := handler.service.CreateEmployee(*employeeModel)

		if err != nil {
			handler.handleError(w, err)
			return
		}

		employeeDto := mappers.EmployeeModelToDTO(employeeCreated)
		respondWithJSON(w, http.StatusOK, "Success", employeeDto)
	}
}

func (handler *EmployeeHandler) UpdateEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			handler.handleError(w, err)
			return
		}

		var employeePatchRequestDTO dto.EmployeePatchRequestDTO
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&employeePatchRequestDTO); err != nil {
			handler.handleError(w, error_management.ErrEmployeeBodyDecoding)
			return
		}

		employeePatchRequestModel := mappers.EmployeePatchRequestDTOToModel(employeePatchRequestDTO)
		employeeUpdated, err := handler.service.UpdateEmployee(id, *employeePatchRequestModel)

		if err != nil {
			handler.handleError(w, err)
			return
		}

		respondWithJSON(w, http.StatusOK, "Success", employeeUpdated)
	}
}

func (handler *EmployeeHandler) DeleteEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employeeIdString := chi.URLParam(r, "id")
		employeeId, err := strconv.Atoi(employeeIdString)
		if err != nil {
			handler.handleError(w, err)
			return
		}

		errorService := handler.service.DeleteEmployee(employeeId)
		if errorService != nil {
			handler.handleError(w, errorService)
			return
		}

		respondWithJSON(w, http.StatusOK, "Empleado eliminado correctamente", "")
	}
}

func (handler *EmployeeHandler) GetReportInboundOrdersByEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.URL.Query().Get("id")
		fmt.Printf("id: %s", idString)

		var employeesDTO []dto.EmployeeReportInboundOrdersDTO
		employees, err := handler.service.GetReportInboundOrdersByEmployee(idString)
		for _, employee := range employees {
			employeeDto := mappers.EmployeeReportInboundOrdersModelToDTO(employee)
			employeesDTO = append(employeesDTO, *employeeDto)
		}

		if err != nil {
			handler.handleError(w, err)
			return
		}

		respondWithJSON(w, http.StatusOK, "Success", employeesDTO)
	}
}

func respondWithJSON[T any](w http.ResponseWriter, code int, msg string, data T) {
	response.JSON(w, code, dto.EmployeeResponseDto[T]{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func (handler *EmployeeHandler) handleError(w http.ResponseWriter, err error) {
	errorEmployee := error_management.HandleErrorEmployee(err)
	respondWithJSON(w, errorEmployee.Code, errorEmployee.Message, (*dto.EmployeeDTO)(nil))
}
