package router

import (
	"context"
	"database/sql"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers"
	"github.com/melisource/fury_go-core/pkg/web"
)

func InitLocalityRouter(rt *web.Router, db *sql.DB, ctx *context.Context) {
	repositoryImp := repository.NewLocalityRepository(db)

	serviceImp := service.NewLocalityService(repositoryImp)

	handler := handlers.NewLocalityHandler(serviceImp)

	group := rt.Group("/api/v1/localities")
	group.Post("/", handler.CreateLocality(ctx))
	group.Get("/reportSellers", handler.GetSellersByLocality(ctx))
	group.Get("/reportCarries", handler.GetCarriersByLocality(ctx))
}
