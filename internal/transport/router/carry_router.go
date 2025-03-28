package router

import (
	"context"
	"database/sql"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers"
	"github.com/melisource/fury_go-core/pkg/web"
)

func InitCarryRouter(rt *web.Router, db *sql.DB, ctx *context.Context) {
	repositoryImp := repository.NewCarrierRepository(db)

	serviceImp := service.NewCarryService(repositoryImp)

	handler := handlers.NewCarryHandler(serviceImp)

	group := rt.Group("/api/v1/carrier")
	group.Get("/", handler.GetAllCarriers(ctx))
	group.Post("/", handler.CreateCarrier(ctx))
}
