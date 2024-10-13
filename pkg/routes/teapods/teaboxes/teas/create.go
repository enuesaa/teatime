package teas

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/srvtea"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	cc := ctx.Use(c)
	teapod := cc.Param("teapodName")
	teabox := cc.Param("teaboxName")

	var reqbody map[string]interface{}
	if err := cc.BindBody(&reqbody); err != nil {
		return err
	}

	teaSrv, err := srvtea.New(cc.Repos, teapod, teabox)
	if err != nil {
		return err
	}

	teaId, err := teaSrv.Create(reqbody)
	if err != nil {
		return err
	}

	res := Creation{
		Id: teaId,
	}
	return cc.WithData(res)
}
