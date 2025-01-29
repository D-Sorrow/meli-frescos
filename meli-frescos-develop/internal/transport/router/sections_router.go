package router

import (
	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/loader"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func InitSectionsRouter(rt *chi.Mux) {
	loader := loader.NewSectionsJSONFile("../docs/db/sections_data.json")

	db, err := loader.Load()
	if err != nil {
		return
	}
	repositoryImp := repository.NewSectionsRepository(db)
	serviceImp := service.NewSectionsService(repositoryImp)

	handler := handlers.NewSectionsHandler(serviceImp)

	rt.Route("/sections", func(rt chi.Router) {
		rt.Get("/", handler.GetSections())
	})
	return

}
