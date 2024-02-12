package main

import (
	"log"

	"github.com/enuesaa/teatime/pkg/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	app.Use(middleware.RecoverWithConfig(middleware.DefaultRecoverConfig))
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	api := app.Group("/api")
	api.GET("/providers", controller.ListProviders)
	api.GET("/providers/:name", controller.DescribeProvider)
	api.POST("/providers", controller.AddProvider)
	api.PUT("/providers/:name", controller.UpdateProvider)
	api.DELETE("/providers/:name", controller.DeleteProvider)
	
	// api.GET("/providers/:name/config", controller.DescribeCard)
	// api.GET("/providers/:name/rows", controller.DescribeCard)
	// api.GET("/providers/:name/rows/:id", controller.DescribeCard)
	// api.POST("/providers/:name/rows", controller.DescribeCard)
	// api.PUT("/providers/:name/rows/:id", controller.DescribeCard)
	// api.DELETE("/providers/:name/rows/:id", controller.DescribeCard)

	if err := app.Start(":3000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
