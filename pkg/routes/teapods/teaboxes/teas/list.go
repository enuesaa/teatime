package teas

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/srvtea"
	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	cc := ctx.Use(c)
	teapod := cc.Param("teapod")
	teabox := cc.Param("teabox")

	teaSrv, err := srvtea.New(cc.Repos, teapod, teabox)
	if err != nil {
		return err
	}
	list, err := teaSrv.List()
	if err != nil {
		return err
	}
	return cc.WithData(list)
}
