package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"

	"github.com/D-Sorrow/meli-frescos/cmd/server"
)

func main() {
	cfg := &server.ConfigServerChi{
		ServerAddress: ":8080",
	}
	cfgDb := &application.ConfigApplicationDefault{
		Db: &mysql.Config{
			User:   "root",
			Passwd: "14074871Gg.",
			Net:    "tcp",
			Addr:   "127.0.0.1:3306",
			DBName: "fantasy_products",
		},
		Addr: "127.0.0.1:8080",
	}
	app := server.NewServerChi(cfg)
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
