package router

import (
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	handler "github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func NewBuyerRouter(rt *chi.Mux, db *sql.DB) {
	buyerRepo := repository.NewBuyerRepository(db)

	buyerService := service.NewBuyerService(buyerRepo)

	buyerHandler := handler.NewBuyerHandler(buyerService)

	rt.Route("/api/v1/buyers", func(rt chi.Router) {
		rt.Get("/", buyerHandler.GetAll())
		rt.Get("/{id}", buyerHandler.GetById())
		rt.Post("/", buyerHandler.Create())
		rt.Patch("/{id}", buyerHandler.Patch())
		rt.Delete("/{id}", buyerHandler.Delete())
		rt.Get("/reportPurchaseOrders", buyerHandler.GetReportPurchaseOrders())
	})
}
