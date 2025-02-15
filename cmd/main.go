package main

import (
	"fmt"
	"log"

	"github.com/D-Sorrow/meli-frescos/cmd/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}
	cfg := &server.ConfigServerChi{
		ServerAddress: ":8080",
	}
	app := server.NewServerChi(cfg)
	if err := app.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
