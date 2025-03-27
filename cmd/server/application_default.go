package server

import (
	"context"

	db_config "github.com/D-Sorrow/meli-frescos/internal/infrastructure/config"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/db"
	"github.com/D-Sorrow/meli-frescos/internal/transport/middlewares"
	"github.com/D-Sorrow/meli-frescos/internal/transport/router"
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

	err = app.Run()

	return
}
