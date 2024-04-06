package controller

import (
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

func GetCard(c echo.Context) error {
	name := c.Param("name")
	id := c.Param("id")

	card, err := service.NewTeapodSrv(name).GetCard(id)
	if err != nil {
		return err
	}

	return WithData(c, card)
}
