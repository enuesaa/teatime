package main

import (
	"log"

	"github.com/enuesaa/teatime/pkg/router"
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	repos := repository.New()
	if err := repos.DB.Connect(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	defer repos.DB.Disconnect()

	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3001"},
	}))

	router.Setup(app, repos)

	if err := app.Start(":3000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
