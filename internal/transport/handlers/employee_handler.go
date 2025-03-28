package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/ports/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/dto"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/error_management"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers/mappers"
	"github.com/melisource/fury_go-core/pkg/web"
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

func (handler *EmployeeHandler) GetEmployees(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		employees, err := handler.service.GetEmployees()

		if err != nil {
			handler.handleError(w, err, ctx)
			return nil
		}
		var employeesDto []dto.EmployeeDTO
		for _, employee := range employees {
			employeeDto := mappers.EmployeeModelToDTO(employee)
			employeesDto = append(employeesDto, *employeeDto)
		}

		respondWithJSON(w, http.StatusOK, "Success", employeesDto)
		return nil
	}
}

func (handler *EmployeeHandler) GetEmployeeById(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		idString := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			handler.handleError(w, err, ctx)
			return nil
		}

		employee, err := handler.service.GetEmployeeById(id)
		employeeDto := mappers.EmployeeModelToDTO(employee)

		if err != nil {
			handler.handleError(w, err, ctx)
			return nil
		}

		respondWithJSON(w, http.StatusOK, "Success", employeeDto)
		return nil
	}
}

func (handler *EmployeeHandler) CreateEmployee(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var employeeToCreate dto.EmployeeRequestDTO
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&employeeToCreate); err != nil {
			handler.handleError(w, error_management.ErrEmployeeBodyDecoding, ctx)
			return nil
		}

		if err := handler.validator.Struct(employeeToCreate); err != nil {
			handler.handleError(w, err, ctx)
			return nil
		}

		employeeModel := mappers.EmployeeDTOToModel(employeeToCreate)
		employeeCreated, err := handler.service.CreateEmployee(*employeeModel)

		if err != nil {
			handler.handleError(w, err, ctx)
			return nil
		}

		employeeDto := mappers.EmployeeModelToDTO(employeeCreated)
		respondWithJSON(w, http.StatusOK, "Success", employeeDto)
		return nil
	}
}

func (handler *EmployeeHandler) UpdateEmployee(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		idString := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			handler.handleError(w, err, ctx)
			return nil
		}

		var employeePatchRequestDTO dto.EmployeePatchRequestDTO
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&employeePatchRequestDTO); err != nil {
			handler.handleError(w, error_management.ErrEmployeeBodyDecoding, ctx)
			return nil
		}

		employeePatchRequestModel := mappers.EmployeePatchRequestDTOToModel(employeePatchRequestDTO)
		employeeUpdated, err := handler.service.UpdateEmployee(id, *employeePatchRequestModel)

		if err != nil {
			handler.handleError(w, err, ctx)
			return nil
		}

		employeeDto := mappers.EmployeeModelToDTO(employeeUpdated)
		respondWithJSON(w, http.StatusOK, "Success", employeeDto)
		return nil
	}
}

func (handler *EmployeeHandler) DeleteEmployee(ctx *context.Context) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		employeeIdString := chi.URLParam(r, "id")
		employeeId, err := strconv.Atoi(employeeIdString)
		if err != nil {
			handler.handleError(w, err, ctx)
			return nil
		}

		errorService := handler.service.DeleteEmployee(employeeId)
		if errorService != nil {
			handler.handleError(w, errorService, ctx)
			return nil
		}

		respondWithJSON(w, http.StatusOK, "Empleado eliminado correctamente", "")
		return nil
	}
}

func (handler *EmployeeHandler) GetReportInboundOrdersByEmployee(
	ctx *context.Context,
) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		idString := r.URL.Query().Get("id")
		fmt.Printf("id: %s", idString)

		var employeesDTO []dto.EmployeeReportInboundOrdersDTO
		employees, err := handler.service.GetReportInboundOrdersByEmployee(idString)
		for _, employee := range employees {
			employeeDto := mappers.EmployeeReportInboundOrdersModelToDTO(employee)
			employeesDTO = append(employeesDTO, *employeeDto)
		}

		if err != nil {
			handler.handleError(w, err, ctx)
			return nil
		}

		respondWithJSON(w, http.StatusOK, "Success", employeesDTO)
		return nil
	}
}

func respondWithJSON[T any](w http.ResponseWriter, code int, msg string, data T) {
	response.JSON(w, code, dto.EmployeeResponseDto[T]{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func (handler *EmployeeHandler) handleError(
	w http.ResponseWriter,
	err error,
	ctx *context.Context,
) {
	errorEmployee := error_management.HandleErrorEmployee(err, ctx)
	respondWithJSON(w, errorEmployee.Code, errorEmployee.Message, (*dto.EmployeeDTO)(nil))
}
