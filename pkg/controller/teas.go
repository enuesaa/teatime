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

	providerSrv := service.NewProviderService("links")
	ids, err := providerSrv.ListTeas()
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

	providerSrv := service.NewProviderService("links")
	row, err := providerSrv.GetTea(rid)
	if err != nil {
		return err
	}

	return WithData(c, row.Value)
}

func CreateTea(c echo.Context) error {
	var value plug.Value
	if err := Validate(c, &value, validationRules); err != nil {
		return err
	}
	providerSrv := service.NewProviderService("links")
	rid, err := providerSrv.CreateTea(value)
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
	providerSrv := service.NewProviderService("links")
	if _, err := providerSrv.UpdateTea(rid, value); err != nil {
		return err
	}
	return WithData(c, NewIdSchema(rid))
}

func DeleteTea(c echo.Context) error {
	rid := c.Param("rid")

	providerSrv := service.NewProviderService("links")
	if err := providerSrv.DeleteTea(rid); err != nil {
		return err
	}

	return WithData(c, NewEmptySchema())
}
