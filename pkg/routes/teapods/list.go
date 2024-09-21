package teapods

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	cc := ctx.Use(c)

	teapodSrv := service.NewTeapodSrv(cc.Repos)
	list, err := teapodSrv.List()
	if err != nil {
		return err
	}
	return cc.WithData(list)
}
