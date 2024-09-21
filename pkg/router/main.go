package router

import (
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/routes/teapods"
	"github.com/enuesaa/teatime/pkg/routes/teapods/teas"
	"github.com/enuesaa/teatime/ui"
	"github.com/labstack/echo/v4"
)

func Setup(app *echo.Echo, repos repository.Repos) {
	api := app.Group("/api")
	api.Use(HandleData)
	api.Use(HandleError)

	api.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
			cc := ctx.Context{
				Context: c,
				Repos: repos,
			}
            return next(cc)
        }
    })

	api.GET("/teapods", teapods.List)
	api.GET("/teapods/:teapod", teapods.View)
	api.GET("/teapods/:teapod/teas", teas.List)

	app.Any("/*", ui.Serve)
}
