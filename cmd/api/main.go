package main

import (
	"log"

	"github.com/melisource/fury_bootcamp-go-w15-s4-3-6/cmd/server"
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
