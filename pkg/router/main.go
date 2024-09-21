package router

import (
	"github.com/enuesaa/teatime/pkg/router/teapods"
	"github.com/enuesaa/teatime/pkg/router/teapods/teas"
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/ui"
	"github.com/labstack/echo/v4"
)

func Setup(app *echo.Echo, repos repository.Repos) {
	api := app.Group("/api")
	api.Use(HandleData)
	api.Use(HandleError)

	teapodsCtl := teapods.New(repos)
	api.GET("/teapods", teapodsCtl.ListTeapods)
	api.GET("/teapods/:teapod", teapodsCtl.GetTeapod)

	teaCtl := teas.New(repos)
	api.GET("/teapods/:teapod/teas", teaCtl.ListTeas)

	app.Any("/*", ui.Serve)
}
