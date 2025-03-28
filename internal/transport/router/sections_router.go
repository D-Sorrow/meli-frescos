package router

import (
	"context"
	"database/sql"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service"
	"github.com/melisource/fury_go-core/pkg/web"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers"
)

func InitSectionsRouter(rt *web.Router, db *sql.DB, ctx *context.Context) {

	repositoryImp := repository.NewSectionsRepository(db)
	serviceImp := service.NewSectionsService(repositoryImp)

	handler := handlers.NewSectionsHandler(serviceImp)

	group := rt.Group("/sections")
	group.Get("/", handler.GetSections(ctx))
	group.Get("/{id}", handler.GetSectionsById(ctx))
	group.Post("/", handler.SaveSections(ctx))
	group.Delete("/{id}", handler.DeleteSections(ctx))
}
