package router

import (
	"context"
	"database/sql"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers"
	"github.com/melisource/fury_go-core/pkg/web"
)

func InitWarehouseRouter(rt *web.Router, db *sql.DB, ctx *context.Context) {
	repositoryImp := repository.NewWarehouseRepository(db)

	serviceImp := service.NewWarehouseService(repositoryImp)

	handler := handlers.NewWarehouseHandler(serviceImp)

	group := rt.Group("/api/v1/warehouses")
	group.Get("/", handler.GetWarehouses(ctx))
	group.Get("/{id}", handler.GetWarehouseById(ctx))
	group.Delete("/{id}", handler.DeleteWarehouse(ctx))
	group.Post("/", handler.CreateWarehouse(ctx))
	group.Patch("/{id}", handler.PatchWarehouse(ctx))
}
