package router

import (
	"context"
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	handler "github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func NewPurchaseOrderRouter(rt *chi.Mux, db *sql.DB, ctx *context.Context) {
	purchaseOrderRepo := repository.NewPurchaseOrderRepository(db)

	purchaseOrderService := service.NewPurchaseOrderService(purchaseOrderRepo)

	purchaseOrderHandler := handler.NewPurchaseOrderHandler(purchaseOrderService)

	rt.Route("/api/v1/purchaseOrders", func(rt chi.Router) {
		rt.Get("/{id}", purchaseOrderHandler.GetById(ctx))
		rt.Post("/", purchaseOrderHandler.Create(ctx))
	})
}
