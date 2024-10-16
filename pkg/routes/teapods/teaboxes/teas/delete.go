package teas

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/srvtea"
	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	cc := ctx.Use(c)
	teapod := cc.Param("teapodName")
	teabox := cc.Param("teaboxName")
	teaId := cc.Param("teaId")

	teaSrv := srvtea.New(cc.Repos, teapod, teabox)
	if err := teaSrv.Delete(teaId); err != nil {
		return err
	}

	return cc.WithData(DeleteResponse{})
}
