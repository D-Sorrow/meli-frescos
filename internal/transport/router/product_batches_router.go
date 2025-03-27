package router

import (
	"context"
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func InitProductBatchesRouter(rt *chi.Mux, db *sql.DB, ctx *context.Context) {

	repositoryImp := repository.NewProductBatchesRepository(db)

	serviceImp := service.NewProductBatches(repositoryImp)

	handler := handlers.NewProductBatches(serviceImp)

	rt.Route("/api/v1/probatch", func(rt chi.Router) {
		rt.Post("/", handler.AddProductBatches(ctx))
		rt.Get("/{id}", handler.GetById(ctx))
		rt.Post("/2", handler.Create(ctx))

	})
}
