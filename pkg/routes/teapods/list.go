package teapods

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/srvteapod"
	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	cc := ctx.Use(c)

	teapodSrv := srvteapod.New(cc.Repos)
	teapodNames, err := teapodSrv.List()
	if err != nil {
		return err
	}

	list := []Item{}
	for _, teapodName := range teapodNames {
		list = append(list, Item{
			Name: teapodName,
		})
	}

	return cc.WithItems(list)
}
