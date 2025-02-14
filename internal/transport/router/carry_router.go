package router

import (
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func InitCarryRouter(rt *chi.Mux, db *sql.DB) {
	repositoryImp := repository.NewCarrierRepository(db)

	serviceImp := service.NewCarryService(repositoryImp)

	handler := handlers.NewCarryHandler(serviceImp)

	rt.Route("/api/v1/carrier", func(rt chi.Router) {
		rt.Get("/", handler.GetAllCarriers())
		rt.Post("/", handler.CreateCarrier())
	})
}
