package main

import (
	"log"

	"github.com/enuesaa/teatime/ui"
	"github.com/enuesaa/teatime/pkg/controller"
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	repos := repository.New()
	// migrate if table does not exist.
	if err := repos.DB.Migrate(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

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

	// api teapods
	api.GET("/teapods", controller.ListTeapods)
	api.GET("/teapods/:name", controller.GetTeapodInfo)
	api.GET("/teapods/:name/cards/:id", controller.GetCard)

	// api teapods teas
	api.GET("/teapods/:name/teas", controller.ListTeas)
	api.GET("/teapods/:name/teas/:teaid", controller.GetTea)
	api.POST("/teapods/:name/teas", controller.CreateTea)
	api.PUT("/teapods/:name/teas/:teaid", controller.UpdateTea)
	api.DELETE("/teapods/:name/teas/:teaid", controller.DeleteTea)

	// ui
	app.Any("/*", ui.Serve)

	if err := app.Start(":3000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
