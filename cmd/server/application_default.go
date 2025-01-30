package server

import (
	"net/http"

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

	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	router.InitSellerRouter(rt)

	router.InitProductRouter(rt)
	router.NewBuyerRouter(rt)

	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
