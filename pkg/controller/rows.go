package controller

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

var validationRules = map[string]interface{}{
	"name":        "required",
	"description": "required",
}

func ListRows(c echo.Context) error {
	list := make([]IdSchema, 0)

	providerSrv := service.NewProviderService("links")
	ids, err := providerSrv.ListRows()
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

func DescribeRow(c echo.Context) error {
	name := c.Param("name")

	providerSrv := service.NewProviderService("links")
	row, err := providerSrv.GetRow(name)
	if err != nil {
		return err
	}

	return WithData(c, row.Value)
}

func CreateRow(c echo.Context) error {
	var values plug.Value
	if err := Validate(c, &values, validationRules); err != nil {
		return err
	}
	providerSrv := service.NewProviderService("links")
	id, err := providerSrv.CreateRow(values)
	if err != nil {
		return err
	}
	return WithData(c, NewIdSchema(id))
}

func UpdateRow(c echo.Context) error {
	id := c.Param("id")

	var values plug.Value
	if err := Validate(c, &values, validationRules); err != nil {
		return err
	}
	providerSrv := service.NewProviderService("links")
	_, err := providerSrv.UpdateRow(id, values)
	if err != nil {
		return err
	}
	return WithData(c, NewIdSchema(id))
}

func DeleteRow(c echo.Context) error {
	id := c.Param("id")

	providerSrv := service.NewProviderService("coredata")
	if err := providerSrv.DeleteRow(id); err != nil {
		return err
	}

	return WithData(c, NewEmptySchema())
}
