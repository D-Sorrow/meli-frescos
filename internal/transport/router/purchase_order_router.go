package router

import (
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	handler "github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func NewPurchaseOrderRouter(rt *chi.Mux, db *sql.DB) {
	purchaseOrderRepo := repository.NewPurchaseOrderRepository(db)

	purchaseOrderService := service.NewPurchaseOrderService(purchaseOrderRepo)

	purchaseOrderHandler := handler.NewPurchaseOrderHandler(purchaseOrderService)

	rt.Route("/api/v1/purchaseOrders", func(rt chi.Router) {
		rt.Get("/{id}", purchaseOrderHandler.GetById())
		rt.Post("/", purchaseOrderHandler.Create())
	})
}
