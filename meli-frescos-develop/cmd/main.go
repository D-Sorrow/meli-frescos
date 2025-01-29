package main

import (
	"fmt"

	"github.com/D-Sorrow/meli-frescos/cmd/server"
)

func main() {
	fmt.Print("masdas")
	cfg := &server.ConfigServerChi{
		ServerAddress:  ":8089",
		LoaderFilePath: "docs/db/vehicles_100.json",
	}
	app := server.NewServerChi(cfg)
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
