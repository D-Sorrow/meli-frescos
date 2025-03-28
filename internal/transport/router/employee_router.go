package router

import (
	"context"
	"database/sql"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers"
	"github.com/melisource/fury_go-core/pkg/web"
)

func InitEmployeeRouter(rt *web.Router, db *sql.DB, ctx *context.Context) {
	repositoryImp := repository.NewEmployeeRepository(db)

	serviceImp := service.NewEmployeeService(repositoryImp)

	handler := handlers.NewEmployeeHandler(serviceImp)

	group := rt.Group("/api/v1/employees")
	group.Get("/", handler.GetEmployees(ctx))
	group.Get("/{id}", handler.GetEmployeeById(ctx))
	group.Post(("/"), handler.CreateEmployee(ctx))
	group.Patch("/{id}", handler.UpdateEmployee(ctx))
	group.Delete("/{id}", handler.DeleteEmployee(ctx))
	group.Get(("/reportinboundorders"), handler.GetReportInboundOrdersByEmployee(ctx))
}
