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
	teapodName := c.Param("teapod")

	teas, err := service.NewTeapodSrv(teapodName).ListTeas()
	if err != nil {
		return err
	}
	list := make([]IdSchema, 0)
	for _, tea := range teas {
		list = append(list, IdSchema{
			Id: tea.Teaid,
		})
	}
	return WithData(c, list)
}

func GetTea(c echo.Context) error {
	teapodName := c.Param("teapod")
	teaid := c.Param("teaid")

	tea, err := service.NewTeapodSrv(teapodName).GetTea(teaid)
	if err != nil {
		return err
	}
	return WithData(c, tea.Value)
}

func CreateTea(c echo.Context) error {
	teapodName := c.Param("teapod")

	var value plug.Value
	if err := Validate(c, &value, validationRules); err != nil {
		return err
	}
	teaid, err := service.NewTeapodSrv(teapodName).CreateTea(value)
	if err != nil {
		return err
	}
	return WithData(c, NewIdSchema(teaid))
}

func UpdateTea(c echo.Context) error {
	teapodName := c.Param("teapod")
	teaid := c.Param("teaid")

	var value plug.Value
	if err := Validate(c, &value, validationRules); err != nil {
		return err
	}
	if _, err := service.NewTeapodSrv(teapodName).UpdateTea(teaid, value); err != nil {
		return err
	}
	return WithData(c, NewIdSchema(teaid))
}

func DeleteTea(c echo.Context) error {
	teapodName := c.Param("teapod")
	teaid := c.Param("teaid")

	if err := service.NewTeapodSrv(teapodName).DeleteTea(teaid); err != nil {
		return err
	}

	return WithData(c, NewEmptySchema())
}
