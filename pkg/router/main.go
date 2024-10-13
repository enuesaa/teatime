package router

import (
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/router/middleware"
	apiTeapods "github.com/enuesaa/teatime/pkg/routes/teapods"
	apiTeapodsInfo "github.com/enuesaa/teatime/pkg/routes/teapods/info"
	apiTeapodsTeaboxesTeas "github.com/enuesaa/teatime/pkg/routes/teapods/teaboxes/teas"
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

	api.GET("/teapods", apiTeapods.List)
	api.POST("/teapods", apiTeapods.Create)
	api.DELETE("/teapods/:teapodName", apiTeapods.Delete)

	api.GET("/teapods/:teapodName/info", apiTeapodsInfo.View)
	api.GET("/teapods/:teapodName/teaboxes/:teaboxName/teas", apiTeapodsTeaboxesTeas.List)
	api.GET("/teapods/:teapodName/teaboxes/:teaboxName/teas/:teaId", apiTeapodsTeaboxesTeas.View)
	api.POST("/teapods/:teapodName/teaboxes/:teaboxName/teas", apiTeapodsTeaboxesTeas.Create)

	app.Any("/*", ui.Serve)

	return app
}
