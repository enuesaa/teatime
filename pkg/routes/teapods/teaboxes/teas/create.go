package teas

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/srvtea"
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

	teaSrv := srvtea.New(cc.Repos, teapodName, teaboxName)
	teaId, err := teaSrv.Create(reqbody)
	if err != nil {
		return err
	}
	res := Creation{
		Id: teaId,
	}
	return cc.WithData(res)
}
