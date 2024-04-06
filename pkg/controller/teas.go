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
	name := c.Param("name")

	ids, err := service.NewTeapodSrv(name).ListTeas()
	if err != nil {
		return err
	}
	list := make([]IdSchema, 0)
	for _, id := range ids {
		list = append(list, IdSchema{
			Id: id,
		})
	}
	return WithData(c, list)
}

func GetTea(c echo.Context) error {
	name := c.Param("name")
	teaid := c.Param("teaid")

	tea, err := service.NewTeapodSrv(name).GetTea(teaid)
	if err != nil {
		return err
	}
	return WithData(c, tea.Value)
}

func CreateTea(c echo.Context) error {
	name := c.Param("name")

	var value plug.Value
	if err := Validate(c, &value, validationRules); err != nil {
		return err
	}
	rid, err := service.NewTeapodSrv(name).CreateTea(value)
	if err != nil {
		return err
	}
	return WithData(c, NewIdSchema(rid))
}

func UpdateTea(c echo.Context) error {
	name := c.Param("name")
	teaid := c.Param("teaid")

	var value plug.Value
	if err := Validate(c, &value, validationRules); err != nil {
		return err
	}
	if _, err := service.NewTeapodSrv(name).UpdateTea(teaid, value); err != nil {
		return err
	}
	return WithData(c, NewIdSchema(teaid))
}

func DeleteTea(c echo.Context) error {
	name := c.Param("name")
	teaid := c.Param("teaid")

	if err := service.NewTeapodSrv(name).DeleteTea(teaid); err != nil {
		return err
	}

	return WithData(c, NewEmptySchema())
}
