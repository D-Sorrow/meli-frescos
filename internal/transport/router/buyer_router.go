package router

import (
	"context"
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	handler "github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func NewBuyerRouter(rt *chi.Mux, db *sql.DB, ctx *context.Context) {
	buyerRepo := repository.NewBuyerRepository(db)

	buyerService := service.NewBuyerService(buyerRepo)

	buyerHandler := handler.NewBuyerHandler(buyerService)

	rt.Route("/api/v1/buyers", func(rt chi.Router) {
		rt.Get("/", buyerHandler.GetAll(ctx))
		rt.Get("/{id}", buyerHandler.GetById(ctx))
		rt.Post("/", buyerHandler.Create(ctx))
		rt.Patch("/{id}", buyerHandler.Patch(ctx))
		rt.Delete("/{id}", buyerHandler.Delete(ctx))
		rt.Get("/reportPurchaseOrders", buyerHandler.GetReportPurchaseOrders(ctx))
	})
}
