package main

import (
	"log"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/router"
)

func main() {
	repos := repository.New()
	if err := repos.Startup(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	defer repos.End()

	app := router.New(repos)

	if err := app.Start(":3000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
