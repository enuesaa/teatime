package main

import (
	"log"

	"github.com/enuesaa/teatime/pkg/controller"
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/ui"
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

	// api
	api := app.Group("/api")
	ctl := controller.New(repos)
	api.Use(ctl.HandleData)
	api.Use(ctl.HandleError)
	api.GET("/teapods", ctl.ListTeapods)
	api.GET("/teapods/:teapod", ctl.GetTeapod)
	api.GET("/teapods/:teapod/:teabox", ctl.ListTeas)

	// ui
	app.Any("/*", ui.Serve)

	if err := app.Start(":3000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
