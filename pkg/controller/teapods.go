package controller

import (
	"github.com/labstack/echo/v4"
)

type Teapod struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Command string `json:"command"`
}
func ListTeapods(c echo.Context) error {
	list := make([]Teapod, 0)
	list = append(list, Teapod{
		Id: "links",
		Name: "links",
		Command: "teapod-links",
	})
	return WithData(c, list)
}
