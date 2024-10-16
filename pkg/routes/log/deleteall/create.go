package deleteall

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/srvlog"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	cc := ctx.Use(c)

	var reqbody CreateRequestBody
	if err := cc.Bind(&reqbody); err != nil {
		return err
	}

	logSrv := srvlog.New(cc.Repos)
	if err := logSrv.DeleteAll(); err != nil {
		return err
	}

	res := Creation{
		Ok: true,
	}
	return cc.WithData(res)
}
