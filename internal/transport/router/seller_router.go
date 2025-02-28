package router

import (
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func InitSellerRouter(rt *chi.Mux, db *sql.DB) {

	repositoryImp := repository.NewSellerRepository(db)

	serviceImp := service.NewSellerService(repositoryImp)

	handler := handlers.NewHandlerService(serviceImp)

	rt.Route("/api/v1/sellers", func(rt chi.Router) {
		rt.Get("/", handler.GetSellers())
		rt.Get("/{id}", handler.GetSeller())
		rt.Post("/", handler.CreateSeller())
		rt.Patch("/{id}", handler.UpdateSeller())
		rt.Delete("/{id}", handler.DeleteSeller())
	})
}
