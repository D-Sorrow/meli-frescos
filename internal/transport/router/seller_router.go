package router

import (
	"context"
	"database/sql"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers"
	"github.com/melisource/fury_go-core/pkg/web"
)

func InitSellerRouter(rt *web.Router, db *sql.DB, ctx *context.Context) {

	repositoryImp := repository.NewSellerRepository(db)

	serviceImp := service.NewSellerService(repositoryImp)

	handler := handlers.NewHandlerService(serviceImp)

	group := rt.Group("/api/v1/sellers")
	group.Get("/", handler.GetSellers(ctx))
	group.Get("/{id}", handler.GetSeller(ctx))
	group.Post("/", handler.CreateSeller(ctx))
	group.Patch("/{id}", handler.UpdateSeller(ctx))
	group.Delete("/{id}", handler.DeleteSeller(ctx))
}
