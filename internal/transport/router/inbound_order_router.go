package router

import (
	"context"
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func InitInboundOrderRouter(rt *chi.Mux, db *sql.DB, ctx *context.Context) {

	repositoryImp := repository.NewInboundOrderRepository(db)

	serviceImp := service.NewInboundOrderService(repositoryImp)

	handler := handlers.NewInboundOrderHandler(serviceImp)

	rt.Route("/api/v1/inboundOrders", func(rt chi.Router) {
		rt.Post(("/"), handler.CreateInboundOrder(ctx))
	})
}
