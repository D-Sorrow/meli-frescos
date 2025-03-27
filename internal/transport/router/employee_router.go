package router

import (
	"context"
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func InitEmployeeRouter(rt *chi.Mux, db *sql.DB, ctx *context.Context) {

	repositoryImp := repository.NewEmployeeRepository(db)

	serviceImp := service.NewEmployeeService(repositoryImp)

	handler := handlers.NewEmployeeHandler(serviceImp)

	rt.Route("/api/v1/employees", func(rt chi.Router) {
		rt.Get("/", handler.GetEmployees(ctx))
		rt.Get("/{id}", handler.GetEmployeeById(ctx))
		rt.Post(("/"), handler.CreateEmployee(ctx))
		rt.Patch("/{id}", handler.UpdateEmployee(ctx))
		rt.Delete("/{id}", handler.DeleteEmployee(ctx))
		rt.Get(("/reportinboundorders"), handler.GetReportInboundOrdersByEmployee(ctx))
	})
}
