package teas

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/srvteapod"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	cc := ctx.Use(c)
	teapodName := cc.Param("teapodName")
	teaboxName := cc.Param("teaboxName")

	var reqbody map[string]interface{}
	if err := cc.BindBody(&reqbody); err != nil {
		return err
	}

	teapodSrv := srvteapod.New(cc.Repos)

	if err := teapodSrv.On(teapodName, plug.EventDataCreated, teaboxName, reqbody); err != nil {
		return err
	}
	res := Creation{
		Id: "", // TODO
	}
	return cc.WithData(res)
}
