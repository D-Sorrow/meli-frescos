package handlers

import (
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
			"data": employee,
		})

	}
}
