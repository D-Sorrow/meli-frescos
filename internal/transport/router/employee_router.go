package router

import (
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func InitEmployeeRouter(rt *chi.Mux, db *sql.DB) {

	repositoryImp := repository.NewEmployeeRepository(db)

	serviceImp := service.NewEmployeeService(repositoryImp)

	handler := handlers.NewEmployeeHandler(serviceImp)

	rt.Route("/api/v1/employees", func(rt chi.Router) {
		rt.Get("/", handler.GetEmployees())
		rt.Get("/{id}", handler.GetEmployeeById())
		rt.Post(("/"), handler.CreateEmployee())
		rt.Patch("/{id}", handler.UpdateEmployee())
		rt.Delete("/{id}", handler.DeleteEmployee())
		rt.Get(("/reportinboundorders"), handler.GetReportInboundOrdersByEmployee())
	})
}
