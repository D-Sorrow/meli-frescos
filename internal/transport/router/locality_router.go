package router

import (
	"context"
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func InitLocalityRouter(rt *chi.Mux, db *sql.DB, ctx *context.Context) {

	repositoryImp := repository.NewLocalityRepository(db)

	serviceImp := service.NewLocalityService(repositoryImp)

	handler := handlers.NewLocalityHandler(serviceImp)

	rt.Route("/api/v1/localities", func(rt chi.Router) {
		rt.Post("/", handler.CreateLocality(ctx))
		rt.Get("/reportSellers", handler.GetSellersByLocality(ctx))
		rt.Get("/reportCarries", handler.GetCarriersByLocality(ctx))
	})
}
