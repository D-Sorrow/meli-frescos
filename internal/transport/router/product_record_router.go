package router

import (
	"context"
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func InitProductRecordRouter(rt *chi.Mux, db *sql.DB, ctx *context.Context) {
	repositoryImp := repository.NewProductRecordRepository(db)

	serviceImp := service.NewProductRecordService(repositoryImp)

	handler := handlers.NewProductRecordHandler(serviceImp)

	rt.Route("/api/v1/productRecords", func(rt chi.Router) {
		rt.Post("/", handler.SaveProductRecord(ctx))
		rt.Get("/", handler.GetProductRecord(ctx))
	})
}
