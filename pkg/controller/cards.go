package controller

import (
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

type Card struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Text string `json:"text"`
}
func GetCard(c echo.Context) error {
	teapodName := c.Param("teapod")
	name := c.Param("name")

	data, err := service.NewTeapodSrv(teapodName).GetCard(name)
	if err != nil {
		return err
	}
	card := Card{
		Name: data.Name,
		Title: data.Title,
		Text: data.Text,
	}

	return WithData(c, card)
}
