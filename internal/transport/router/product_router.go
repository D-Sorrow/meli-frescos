package router

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/loader"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func InitProductRouter(rt *chi.Mux) {
	loader := loader.NewProductJSONFile("docs/db/products_data.json")

	db, err := loader.Load()
	if err != nil {
		return
	}

	repositoryImp := repository.NewProductRepository(db)

	serviceImp := service.NewProductService(repositoryImp)

	handler := handlers.NewProductHandler(serviceImp)

	rt.Route("/api/v1/products", func(rt chi.Router) {
		rt.Get("/", handler.GetProducts())
		rt.Get("/{id}", handler.GetProductByID())
		rt.Post("/", handler.SaveProduct())
		rt.Patch("/{id}", handler.UpdateProduct())
	})
	return
}
