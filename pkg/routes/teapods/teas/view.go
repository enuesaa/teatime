package teas

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/srvteapod"
	"github.com/labstack/echo/v4"
)

func View(c echo.Context) error {
	cc := ctx.Use(c)
	teapod := cc.Param("teapod")
	teaId := cc.Param("teaId")

	teapodSrv, err := srvteapod.New(cc.Repos, teapod)
	if err != nil {
		return err
	}
	data, err := teapodSrv.Get(teaId)
	if err != nil {
		return err
	}
	return cc.WithData(data)
}
