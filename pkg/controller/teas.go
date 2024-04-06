package controller

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

var validationRules = map[string]interface{}{
	// "name":        "required",
	// "description": "required",
}

func ListTeas(c echo.Context) error {
	list := make([]IdSchema, 0)

	ids, err := service.NewTeapodSrv("links").ListTeas()
	if err != nil {
		return err
	}
	for _, id := range ids {
		list = append(list, IdSchema{
			Id: id,
		})
	}

	return WithData(c, list)
}

func GetTea(c echo.Context) error {
	rid := c.Param("rid")

	tea, err := service.NewTeapodSrv("links").GetTea(rid)
	if err != nil {
		return err
	}
	return WithData(c, tea.Value)
}

func CreateTea(c echo.Context) error {
	var value plug.Value
	if err := Validate(c, &value, validationRules); err != nil {
		return err
	}
	rid, err := service.NewTeapodSrv("links").CreateTea(value)
	if err != nil {
		return err
	}
	return WithData(c, NewIdSchema(rid))
}

func UpdateTea(c echo.Context) error {
	rid := c.Param("rid")

	var value plug.Value
	if err := Validate(c, &value, validationRules); err != nil {
		return err
	}
	teapodSrv := service.NewTeapodSrv("links")
	if _, err := teapodSrv.UpdateTea(rid, value); err != nil {
		return err
	}
	return WithData(c, NewIdSchema(rid))
}

func DeleteTea(c echo.Context) error {
	rid := c.Param("rid")

	teapodSrv := service.NewTeapodSrv("links")
	if err := teapodSrv.DeleteTea(rid); err != nil {
		return err
	}

	return WithData(c, NewEmptySchema())
}
