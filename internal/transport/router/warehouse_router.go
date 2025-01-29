package router

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/loader"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func InitWarehouseRouter(rt *chi.Mux) {

	loader := loader.NewWarehouseJSONFile("docs/db/warehouse_data.json")
	db, err := loader.Load()
	if err != nil {
		println(err.Error())
		return
	}

	repositoryImp := repository.NewWarehouseRepository(db)

	serviceImp := service.NewWarehouseService(repositoryImp)

	handler := handlers.NewWarehouseHandler(serviceImp)

	rt.Route("/warehouse", func(rt chi.Router) {
		rt.Get("/", handler.GetWarehouses())
	})
}
