package controller

import (
	"github.com/enuesaa/teatime/pkg/service"
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
