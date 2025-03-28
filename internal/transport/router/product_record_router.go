package router

import (
	"context"
	"database/sql"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/domain/service"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/repository"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/handlers"
	"github.com/melisource/fury_go-core/pkg/web"
)

func InitProductRecordRouter(rt *web.Router, db *sql.DB, ctx *context.Context) {
	repositoryImp := repository.NewProductRecordRepository(db)

	serviceImp := service.NewProductRecordService(repositoryImp)

	handler := handlers.NewProductRecordHandler(serviceImp)

	group := rt.Group("/api/v1/productRecords")
	group.Post("/", handler.SaveProductRecord(ctx))
	group.Get("/", handler.GetProductRecord(ctx))
}
