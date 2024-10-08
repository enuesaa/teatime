package teapods

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/srvteapod"
	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	cc := ctx.Use(c)
	teapodName := cc.Param("teapod")

	teapodSrv := srvteapod.New(cc.Repos)
	if err := teapodSrv.UnRegister(teapodName); err != nil {
		return err
	}

	res := Creation{
		Ok: true,
	}
	return cc.WithData(res)
}
