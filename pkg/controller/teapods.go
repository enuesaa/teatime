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
}
func GetTeapodInfo(c echo.Context) error {
	name := c.Param("name")

	providerSrv := service.NewProviderService(name)
	info, err := providerSrv.GetInfo()
	if err != nil {
		return err
	}

	data := TeapodInfo {
		Name: name,
		Command: fmt.Sprintf("teapod-%s", name),
		Description: info.Description,
	}

	return WithData(c, data)
}

type Schema struct {
	Name string `json:"name"`
	Vals map[string]string `json:"vals"`
}
func ListSchemas(c echo.Context) error {
	providerSrv := service.NewProviderService("links")
	info, err := providerSrv.GetInfo()
	if err != nil {
		return err
	}

	list := make([]Schema, 0)
	for _, schema := range info.Schemas {
		list = append(list, Schema{
			Name: schema.Name,
			Vals: schema.Vals,
		})
	}
	return WithData(c, list)
}
