package act

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/srvteapod"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	cc := ctx.Use(c)
	teapodName := cc.Param("teapodName")

	var reqbody CreateRequestBody
	if err := cc.Bind(&reqbody); err != nil {
		return err
	}

	teapodSrv := srvteapod.New(cc.Repos)
	message, err := teapodSrv.Act(teapodName, reqbody.Action)
	if err != nil {
		return err
	}

	res := Creation{
		Message: message,
	}
	return cc.WithData(res)
}
