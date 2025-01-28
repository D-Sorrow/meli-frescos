package server

import (
	"github.com/D-Sorrow/meli-frescos/internal/transport/router"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ConfigServerChi struct {
	ServerAddress  string
	LoaderFilePath string
}

func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg.ServerAddress == "" {

	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
	}

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

type ServerChi struct {
	serverAddress  string
	loaderFilePath string
}

func (a *ServerChi) Run() (err error) {

	rt := chi.NewRouter()

	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	router.InitSellerRouter(rt)

	router.InitProductRouter(rt)

	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
