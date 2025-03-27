package router

import (
	"context"
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func InitProductRouter(rt *chi.Mux, db *sql.DB, ctx *context.Context) {
	repositoryImp := repository.NewProductRepository(db)

	serviceImp := service.NewProductService(repositoryImp)

	handler := handlers.NewProductHandler(serviceImp)

	rt.Route("/api/v1/products", func(rt chi.Router) {
		rt.Get("/", handler.GetProducts(ctx))
		rt.Get("/{id}", handler.GetProductByID(ctx))
		rt.Post("/", handler.SaveProduct(ctx))
		rt.Patch("/{id}", handler.UpdateProduct(ctx))
		rt.Delete("/{id}", handler.DeleteProduct(ctx))
	})
}
