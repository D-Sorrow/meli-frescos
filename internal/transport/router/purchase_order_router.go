package router

import (
	"context"
	"database/sql"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository"
	handler "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers"
	"github.com/melisource/fury_go-core/pkg/web"
)

func NewPurchaseOrderRouter(rt *web.Router, db *sql.DB, ctx *context.Context) {
	purchaseOrderRepo := repository.NewPurchaseOrderRepository(db)

	purchaseOrderService := service.NewPurchaseOrderService(purchaseOrderRepo)

	purchaseOrderHandler := handler.NewPurchaseOrderHandler(purchaseOrderService)

	group := rt.Group("/api/v1/purchaseOrders")
	group.Get("/{id}", purchaseOrderHandler.GetById(ctx))
	group.Post("/", purchaseOrderHandler.Create(ctx))
}
