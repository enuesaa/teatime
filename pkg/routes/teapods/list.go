package teapods

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/srvteapod"
	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	cc := ctx.Use(c)

	teapodSrv := srvteapod.New(cc.Repos)
	teapods, err := teapodSrv.List()
	if err != nil {
		return err
	}

	list := []Item{}
	for _, teapod := range teapods {
		list = append(list, Item{
			Name: teapod,
		})
	}
	return cc.WithData(list)
}
