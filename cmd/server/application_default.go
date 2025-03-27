package server

import (
	"context"
	"net/http"

	db_config "github.com/D-Sorrow/meli-frescos/internal/infrastructure/config"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/db"
	"github.com/D-Sorrow/meli-frescos/internal/transport/middlewares"
	"github.com/D-Sorrow/meli-frescos/internal/transport/router"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ConfigServerChi struct {
	ServerAddress string
}

func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg == nil {
		cfg = defaultConfig
	} else if cfg.ServerAddress == "" {
		cfg.ServerAddress = defaultConfig.ServerAddress
	}

	return &ServerChi{
		serverAddress: cfg.ServerAddress,
	}
}

type ServerChi struct {
	serverAddress string
}

func (a *ServerChi) Run() (err error) {
	rt := chi.NewRouter()
	ctx := context.Background()
	dbconf, err := db_config.NewConfig()

	if err != nil {
		return
	}

	database := db.NewDataBase(dbconf)

	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)
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
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
