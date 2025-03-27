package main

import (
	"log"

	"github.com/D-Sorrow/meli-frescos/cmd/server"
	"github.com/melisource/fury_go-platform/pkg/fury"
)

func main() {
	app, err := fury.NewWebApplication()
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := server.Run(app); err != nil {
		log.Fatal(err)
		return
	}
}
