package controller

import (
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

func GetCard(c echo.Context) error {
	teapodName := c.Param("teapod")
	id := c.Param("id")

	card, err := service.NewTeapodSrv(teapodName).GetCard(id)
	if err != nil {
		return err
	}

	return WithData(c, card)
}
