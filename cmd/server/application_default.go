package server

import (
	"net/http"

	db_config "github.com/D-Sorrow/meli-frescos/internal/infrastructure/config"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/db"
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
	dbconf, err := db_config.NewConfig()
	database := db.NewDataBase(dbconf)

	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	router.InitWarehouseRouter(rt)
	router.InitSellerRouter(rt)
	router.InitEmployeeRouter(rt, database.Db)
	router.InitInboundOrderRouter(rt, database.Db)

	router.InitProductRouter(rt, database.Db)
	router.NewBuyerRouter(rt)

	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
