package router

import (
	"database/sql"

	"github.com/D-Sorrow/meli-frescos/internal/domain/service"

	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
	"github.com/D-Sorrow/meli-frescos/internal/transport/handlers"
	"github.com/go-chi/chi/v5"
)

func InitSectionsRouter(rt *chi.Mux, db *sql.DB) {

	repositoryImp := repository.NewSectionsRepository(db)
	serviceImp := service.NewSectionsService(repositoryImp)

	handler := handlers.NewSectionsHandler(serviceImp)

	rt.Route("/sections", func(rt chi.Router) {
		rt.Get("/", handler.GetSections())
		rt.Get("/{id}", handler.GetSectionsById())
		rt.Post("/", handler.SaveSections())
		rt.Delete("/{id}", handler.DeleteSections())
		//rt.Patch("/id", handler.UpdateSections())

	})

}
