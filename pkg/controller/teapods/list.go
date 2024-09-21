package teapods

import (
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

func (ctl *Ctl) ListTeapods(c echo.Context) error {
	teapodSrv := service.NewTeapodSrv(ctl.repos)
	list, err := teapodSrv.List()
	if err != nil {
		return err
	}
	return ctl.WithData(c, list)
}
