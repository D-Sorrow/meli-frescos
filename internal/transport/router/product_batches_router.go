package router

import (
	"context"
	"database/sql"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers"
	"github.com/melisource/fury_go-core/pkg/web"
)

func InitProductBatchesRouter(rt *web.Router, db *sql.DB, ctx *context.Context) {
	repositoryImp := repository.NewProductBatchesRepository(db)

	serviceImp := service.NewProductBatches(repositoryImp)

	handler := handlers.NewProductBatches(serviceImp)

	group := rt.Group("/api/v1/probatch")
	group.Post("/", handler.AddProductBatches(ctx))
	group.Get("/{id}", handler.GetById(ctx))
	group.Post("/2", handler.Create(ctx))
}
