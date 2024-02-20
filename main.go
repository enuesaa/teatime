package main

import (
	"fmt"
	"log"

	"github.com/enuesaa/teatime/pkg/controller"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

func main() {
	provider := service.NewProviderService("coredata")
	if err := provider.Init(); err != nil {
		// this is correct. do not stop here to show error message on web console.
		fmt.Printf("Error: %s\n", err.Error())
	}

	app := echo.New()
	app.Use(controller.HandleData)
	app.Use(controller.HandleError)

	api := app.Group("/api")
	api.GET("/rows", controller.ListRows)
	api.GET("/rows/:id", controller.DescribeRow)
	api.POST("/rows", controller.CreateRow)
	api.PUT("/rows/:id", controller.UpdateRow)
	api.DELETE("/rows/:id", controller.DeleteRow)

	if err := app.Start(":3000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
