package main

import (
	"log"

	"github.com/enuesaa/teatime/pkg/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.GET("/providers", controller.ListProviders)
	app.POST("/providers", controller.AddProvider)
	app.GET("/providers/:id", controller.DescribeProvider)
	app.PUT("/providers/:id", controller.UpdateProvider)
	app.DELETE("/providers/:id", controller.DeleteProvider)
	app.GET("/providers/:id/cards/:cardName", controller.DescribeCard)
	app.GET("/providers/:id/panels/:panelName", controller.DescribePanel)
	app.GET("/providers/:id/models/:model/records/:recordName", controller.GetRecord)
	app.POST("/providers/:id/models/:model/records/:recordName", controller.RegisterRecord)
	app.PUT("/providers/:id/models/:model/records/:recordName", controller.SetRecord)
	app.DELETE("/providers/:id/models/:model/records/:recordName", controller.DelRecord)

	if err := app.Start(":3000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
