package main

import (
	"fmt"
	"log"

	"github.com/enuesaa/teatime/frontend"
	"github.com/enuesaa/teatime/pkg/controller"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	provider := service.NewProviderService("links")
	if err := provider.Init(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	return

	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	api := app.Group("/api")
	api.Use(controller.HandleData)
	api.Use(controller.HandleError)
	api.GET("/rows", controller.ListRows)
	api.GET("/rows/:id", controller.DescribeRow)
	api.POST("/rows", controller.CreateRow)
	api.PUT("/rows/:id", controller.UpdateRow)
	api.DELETE("/rows/:id", controller.DeleteRow)

	app.Any("/*", frontend.Serve)

	if err := app.Start(":3000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
