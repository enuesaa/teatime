package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/enuesaa/teatime/pkg/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}
	for _, ctr := range containers {
		fmt.Printf("%s %s (status: %s)\n", ctr.ID, ctr.Image, ctr.Status)
	}

	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
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

	app.Logger.Fatal(app.Start(":3000"))
}
