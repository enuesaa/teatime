package teapods

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	cc := ctx.Use(c)

	teapodSrv := service.NewTeapodSrv(cc.Repos)
	teapods, err := teapodSrv.List()
	if err != nil {
		return err
	}

	list := []Item{}
	for _, teapod := range teapods {
		list = append(list, Item{
			Name: teapod.Name,
		})
	}
	return cc.WithData(list)
}
