package teas

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

func View(c echo.Context) error {
	cc := ctx.Use(c)
	teapod := cc.Param("teapod")
	teaId := cc.Param("teaId")

	teaSrv, err := service.NewTeaSrv(cc.Repos, teapod)
	if err != nil {
		return err
	}
	data, err := teaSrv.Get(teaId)
	if err != nil {
		return err
	}
	return cc.WithData(data)
}
