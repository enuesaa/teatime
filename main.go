package main

import (
	"log"

	"github.com/enuesaa/teatime/pkg/controller"
	// "github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/ui"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// repos := repository.New()

	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3001"},
	}))

	// api
	api := app.Group("/api")
	api.Use(controller.HandleData)
	api.Use(controller.HandleError)

	// api teapod
	teapods := api.Group("/teapods")
	teapods.GET("", controller.ListTeapods)
	teapods.GET("/:teapod", controller.GetTeapodInfo)
	teapods.GET("/:teapod/teas", controller.ListTeas)

	// ui
	app.Any("/*", ui.Serve)

	if err := app.Start(":3000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
