package main

import (
	"log"

	"github.com/enuesaa/teatime/frontend"
	"github.com/enuesaa/teatime/pkg/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	api := app.Group("/api")
	api.Use(controller.HandleData)
	api.Use(controller.HandleError)
	api.GET("/teas", controller.ListTeas)
	api.GET("/teas/:rid", controller.GetTea)
	api.POST("/teas", controller.CreateTea)
	api.PUT("/teas/:rid", controller.UpdateTea)
	api.DELETE("/teas/:rid", controller.DeleteTea)
	// api.GET("/cards", controller.ListCards)
	// api.GET("/cards/:name", controller.GetCard)

	app.Any("/*", frontend.Serve)

	if err := app.Start(":3000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
