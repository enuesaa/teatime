package main

import (
	"fmt"
	"log"

	"github.com/enuesaa/teatime/pkg/controller"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	provider := service.NewProviderService("coredata")
	if err := provider.Init(); err != nil {
		// this is correct. do not stop here to show error message on web console.
		fmt.Printf("Error: %s\n", err.Error())
	}

	app := echo.New()

	app.Use(middleware.RecoverWithConfig(middleware.DefaultRecoverConfig))
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	api := app.Group("/api")
	api.GET("/providers", controller.ListProviders)
	api.GET("/providers/:name", controller.DescribeProvider)
	api.GET("/providers/:name/config", controller.DescribeProviderConfig)

	api.GET("/providers/:name/rows", controller.ListProviderRows)
	api.GET("/providers/:name/rows/:id", controller.DescribeProviderRow)
	api.POST("/providers/:name/rows", controller.CreateProviderRow)
	api.PUT("/providers/:name/rows/:id", controller.UpdateProviderRow)
	api.DELETE("/providers/:name/rows/:id", controller.DeleteProviderRow)

	if err := app.Start(":3000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
