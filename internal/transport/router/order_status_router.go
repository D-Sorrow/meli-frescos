package router

import (
	"context"
	"database/sql"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository"
	handler "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers"
	"github.com/melisource/fury_go-core/pkg/web"
)

func NewOrderStatusRouter(rt *web.Router, db *sql.DB, ctx *context.Context) {
	orderStatusRepo := repository.NewOrderStatusRepository(db)

	orderStatusService := service.NewOrderStatusService(orderStatusRepo)

	orderStatusHandler := handler.NewOrderStatusHandler(orderStatusService)

	group := rt.Group("/api/v1/orderStatus")
	group.Get("/", orderStatusHandler.GetAll(ctx))
}
