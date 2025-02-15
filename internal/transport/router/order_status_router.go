package router

import (
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	handler "github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func NewOrderStatusRouter(rt *chi.Mux, db *sql.DB) {
	orderStatusRepo := repository.NewOrderStatusRepository(db)

	orderStatusService := service.NewOrderStatusService(orderStatusRepo)

	orderStatusHandler := handler.NewOrderStatusHandler(orderStatusService)

	rt.Route("/api/v1/orderStatus", func(rt chi.Router) {
		rt.Get("/", orderStatusHandler.GetAll())
	})
}
