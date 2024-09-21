package info

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

func View(c echo.Context) error {
	cc := ctx.Use(c)
	teapod := cc.Param("teapod")

	teapodSrv := service.NewTeapodSrv(cc.Repos)
	info, err := teapodSrv.GetInfo(teapod)
	if err != nil {
		return err
	}

	data := Item{
		Name:        teapod,
		Command:     fmt.Sprintf("teapod-%s", teapod),
		Description: info.Description,
		Teaboxes:    []ItemTeabox{},
	}

	return cc.WithData(data)
}
