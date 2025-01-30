package router

import (
	"fmt"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/loader"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	handler "github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func NewBuyerRouter(rt *chi.Mux) {

	buyerMap := loader.NewBuyerJSONFile("../docs/db/buyers_100.json")
	db, err := buyerMap.Load()
	if err != nil {
		fmt.Printf("Error loading: %s", err.Error())
		return
	}

	buyerRepo := repository.NewBuyerRepository(db)

	buyerService := service.NewBuyerService(buyerRepo)

	buyerHandler := handler.NewBuyerHandler(buyerService)

	rt.Route("/api/v1/buyers", func(rt chi.Router) {
		rt.Get("/", buyerHandler.GetAll())
		rt.Get("/{id}", buyerHandler.GetById())
		rt.Post("/", buyerHandler.Create())
		rt.Patch("/{id}", buyerHandler.Patch())
		rt.Delete("/{id}", buyerHandler.Delete())
	})
}
