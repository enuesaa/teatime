package main

import (
	"log"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/router"
)

func main() {
	poc()

	repos := repository.New()
	if err := repos.DB.Connect(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	defer repos.DB.Disconnect()

	app := router.New(repos)

	if err := app.Start(":3000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
