package server

import (
	"context"

	db_config "github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/config"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/infrastructure/db"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/middlewares"
	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/internal/transport/router"
	"github.com/melisource/fury_go-platform/pkg/fury"
)

func Run(app *fury.Application) (err error) {
	rt := app.Router
	ctx := context.Background()
	dbconf, err := db_config.NewConfig()

	if err != nil {
		return
	}

	database := db.NewDataBase(dbconf)

	rt.Use(middlewares.LogErrorMiddleware(database, &ctx))

	router.NewBuyerRouter(rt, database.Db, &ctx)
	router.NewPurchaseOrderRouter(rt, database.Db, &ctx)
	router.NewOrderStatusRouter(rt, database.Db, &ctx)
	router.InitLocalityRouter(rt, database.Db, &ctx)
	router.InitSellerRouter(rt, database.Db, &ctx)
	router.InitWarehouseRouter(rt, database.Db, &ctx)
	router.InitEmployeeRouter(rt, database.Db, &ctx)
	router.InitInboundOrderRouter(rt, database.Db, &ctx)
	router.InitProductBatchesRouter(rt, database.Db, &ctx)
	router.InitSectionsRouter(rt, database.Db, &ctx)
	router.InitProductRouter(rt, database.Db, &ctx)
	router.InitProductRecordRouter(rt, database.Db, &ctx)
	router.InitCarryRouter(rt, database.Db, &ctx)

	err = app.Run()

	return
}
