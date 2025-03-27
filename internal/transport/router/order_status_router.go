package router

import (
	"context"
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	handler "github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/melisource/fury_go-core/pkg/web"
)

func NewOrderStatusRouter(rt *web.Router, db *sql.DB, ctx *context.Context) {
	orderStatusRepo := repository.NewOrderStatusRepository(db)

	orderStatusService := service.NewOrderStatusService(orderStatusRepo)

	orderStatusHandler := handler.NewOrderStatusHandler(orderStatusService)

	group := rt.Group("/api/v1/orderStatus")
	group.Get("/", orderStatusHandler.GetAll(ctx))
}
