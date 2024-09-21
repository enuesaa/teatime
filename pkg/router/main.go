package router

import (
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/router/middleware"
	"github.com/enuesaa/teatime/pkg/routes/teapods"
	"github.com/enuesaa/teatime/pkg/routes/teapods/teas"
	"github.com/enuesaa/teatime/ui"
	"github.com/labstack/echo/v4"
)

func New(repos repository.Repos) *echo.Echo {
	app := echo.New()
	app.Use(middleware.BindCtx(repos))
	app.Use(middleware.Logger)
	app.Use(middleware.Cors)

	api := app.Group("/api")
	api.Use(middleware.HandleData)
	api.Use(middleware.HandleError)

	api.GET("/teapods", teapods.List)
	api.GET("/teapods/:teapod", teapods.View)
	api.GET("/teapods/:teapod/teas", teas.List)

	app.Any("/*", ui.Serve)

	return app
}
