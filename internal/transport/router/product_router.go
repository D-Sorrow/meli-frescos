package router

import (
	"context"
	"database/sql"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers"
	"github.com/melisource/fury_go-core/pkg/web"
)

func InitProductRouter(rt *web.Router, db *sql.DB, ctx *context.Context) {
	repositoryImp := repository.NewProductRepository(db)

	serviceImp := service.NewProductService(repositoryImp)

	handler := handlers.NewProductHandler(serviceImp)

	group := rt.Group("/api/v1/products")
	group.Get("/", handler.GetProducts(ctx))
	group.Get("/{id}", handler.GetProductByID(ctx))
	group.Post("/", handler.SaveProduct(ctx))
	group.Patch("/{id}", handler.UpdateProduct(ctx))
	group.Delete("/{id}", handler.DeleteProduct(ctx))
}
