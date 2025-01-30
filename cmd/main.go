package main

import (
	"fmt"

	"github.com/D-Sorrow/meli-frescos/cmd/server"
)

func main() {
	cfg := &server.ConfigServerChi{
		ServerAddress: ":8080",
	}
	app := server.NewServerChi(cfg)
	if err := app.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
