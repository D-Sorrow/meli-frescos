package router

import (
	"context"
	"database/sql"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository"
	handler "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers"
	"github.com/melisource/fury_go-core/pkg/web"
)

func NewBuyerRouter(rt *web.Router, db *sql.DB, ctx *context.Context) {
	buyerRepo := repository.NewBuyerRepository(db)

	buyerService := service.NewBuyerService(buyerRepo)

	buyerHandler := handler.NewBuyerHandler(buyerService)

	group := rt.Group("/api/v1/buyers")
	group.Get("/", buyerHandler.GetAll(ctx))
	group.Get("/{id}", buyerHandler.GetById(ctx))
	group.Post("/", buyerHandler.Create(ctx))
	group.Patch("/{id}", buyerHandler.Patch(ctx))
	group.Delete("/{id}", buyerHandler.Delete(ctx))
	group.Get("/reportPurchaseOrders", buyerHandler.GetReportPurchaseOrders(ctx))
}
