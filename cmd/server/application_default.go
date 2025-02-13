package server

import (
	"errors"
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
	if err != nil {
		return errors.New("Error en la configuracion de la base de datos")
	}
	database := db.NewDataBase(dbconf)

	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	router.InitWarehouseRouter(rt)
	router.InitSellerRouter(rt, database.Db)
	router.InitEmployeeRouter(rt)

	router.InitProductRouter(rt, database.Db)
	router.NewBuyerRouter(rt)

	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
