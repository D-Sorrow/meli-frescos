package router

import (
	"github.com/D-Sorrow/meli-frescos/internal/repository"
	"github.com/D-Sorrow/meli-frescos/internal/service"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/D-Sorrow/meli-frescos/pkg/models"
	"github.com/go-chi/chi/v5"
)

func InitSellerRouter(rt *chi.Mux) {

	mapSeller := make(map[int]models.Seller)
	repositoryImp := repository.NewSellerRepository(mapSeller)

	serviceImp := service.NewSellerService(repositoryImp)

	handler := handlers.NewHandlerService(serviceImp)

	rt.Route("/seller", func(rt chi.Router) {
		rt.Get("/", handler.GetSeller())
	})
	return
}
