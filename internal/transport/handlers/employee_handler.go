package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/D-Sorrow/meli-frescos/internal/domain/ports/service"
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
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		bytes, err := json.Marshal(employees)
		if err != nil {
			// default error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// set header (before code due to it sets by default "text/plain")
		w.Header().Set("Content-Type", "application/json")

		// set status code
		w.WriteHeader(http.StatusOK)

		// write body
		w.Write(bytes)
	}
}
