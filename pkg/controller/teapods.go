package controller

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

type Teapod struct {
	Name string `json:"name"`
	Command string `json:"command"`
}
func ListTeapods(c echo.Context) error {
	list := make([]Teapod, 0)
	list = append(list, Teapod{
		Name: "links",
		Command: fmt.Sprintf("teapod-%s", "links"),
	})
	return WithData(c, list)
}

type TeapodInfo struct {
	Name string `json:"name"`
	Command string `json:"command"`
	Description string `json:"description"`
	Cards []string `json:"cards"`
	Schemas []TeapodInfoSchema `json:"schemas"`
}
type TeapodInfoSchema struct {
	Name string `json:"name"`
	Vals map[string]string `json:"vals"`
}
func GetTeapodInfo(c echo.Context) error {
	teapodName := c.Param("teapod")

	info, err := service.NewTeapodSrv(teapodName).GetInfo()
	if err != nil {
		return err
	}

	data := TeapodInfo {
		Name: teapodName,
		Command: fmt.Sprintf("teapod-%s", teapodName),
		Description: info.Description,
		Cards: info.Cards,
		Schemas: make([]TeapodInfoSchema, 0),
	}
	for _, schema := range info.Schemas {
		data.Schemas = append(data.Schemas, TeapodInfoSchema{
			Name: schema.Name,
			Vals: schema.Vals,
		})
	}

	return WithData(c, data)
}
