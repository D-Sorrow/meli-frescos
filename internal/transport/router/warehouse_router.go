package router

import (
	"context"
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	// "github.com/D-Sorrow/meli-frescos/internal/infrastructure/loader"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func InitWarehouseRouter(rt *chi.Mux, db *sql.DB, ctx *context.Context) {

	// Load data from json file
	// loader := loader.NewWarehouseJSONFile("docs/db/warehouse_data.json")
	// db, err := loader.Load()
	// if err != nil {
	// 	println(err.Error())
	// 	return
	// }

	repositoryImp := repository.NewWarehouseRepository(db)

	serviceImp := service.NewWarehouseService(repositoryImp)

	handler := handlers.NewWarehouseHandler(serviceImp)

	rt.Route("/api/v1/warehouses", func(rt chi.Router) {
		rt.Get("/", handler.GetWarehouses(ctx))
		rt.Get("/{id}", handler.GetWarehouseById(ctx))
		rt.Delete("/{id}", handler.DeleteWarehouse(ctx))
		rt.Post("/", handler.CreateWarehouse(ctx))
		rt.Patch("/{id}", handler.PatchWarehouse(ctx))
	})
}
