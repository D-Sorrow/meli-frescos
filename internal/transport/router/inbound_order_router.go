package router

import (
	"context"
	"database/sql"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers"
	"github.com/melisource/fury_go-core/pkg/web"
)

func InitInboundOrderRouter(rt *web.Router, db *sql.DB, ctx *context.Context) {
	repositoryImp := repository.NewInboundOrderRepository(db)

	serviceImp := service.NewInboundOrderService(repositoryImp)

	handler := handlers.NewInboundOrderHandler(serviceImp)

	group := rt.Group("/api/v1/inboundOrders")
	group.Post(("/"), handler.CreateInboundOrder(ctx))
}
