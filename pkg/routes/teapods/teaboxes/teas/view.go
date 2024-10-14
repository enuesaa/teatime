package teas

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/srvtea"
	"github.com/labstack/echo/v4"
)

func View(c echo.Context) error {
	cc := ctx.Use(c)
	teapod := cc.Param("teapodName")
	teabox := cc.Param("teaboxName")
	teaId := cc.Param("teaId")

	teaSrv := srvtea.New(cc.Repos, teapod, teabox)
	data, err := teaSrv.Get(teaId)
	if err != nil {
		return err
	}

	item := Item{
		Id: data.Id(),
		Data: data.Data,
	}
	return cc.WithData(item)
}
