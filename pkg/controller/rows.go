package controller

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

var validationRules = map[string]interface{}{
	"name": "required",
	"description": "required",
}
func ListRows(c echo.Context) error {
	res := NewListResponse[IdSchema]()

	providerSrv := service.NewProviderService("coredata")
	ids, err := providerSrv.ListRows()
	if err != nil {
		return err
	}
	for _, id := range ids {
		res.Items = append(res.Items, IdSchema{
			Id: id,
		})
	}

	return c.JSON(200, res)
}

func DescribeRow(c echo.Context) error {
    id := c.Param("id")

	providerSrv := service.NewProviderService("coredata")
	row, err := providerSrv.GetRow(id)
	if err != nil {
		return err
	}
	res := NewDescribeResponse[map[string]string](row.Values)
	return c.JSON(200, res)
}

func CreateRow(c echo.Context) error {
	var values plug.Values
	if err := ValidateMap(c, &values, validationRules); err != nil {
		return err
	}
	providerSrv := service.NewProviderService("coredata")
	id, err := providerSrv.CreateRow(values)
	if err != nil {
		return err
	}
	return c.JSON(200, NewIdSchema(id))
}

func UpdateRow(c echo.Context) error {
    id := c.Param("id")

	var values plug.Values
	if err := ValidateMap(c, &values, validationRules); err != nil {
		return err
	}
	providerSrv := service.NewProviderService("coredata")
	_, err := providerSrv.UpdateRow(id, values)
	if err != nil {
		return err
	}
	return c.JSON(200, NewIdSchema(id))
}

func DeleteRow(c echo.Context) error {
	id := c.Param("id")

	providerSrv := service.NewProviderService("coredata")
	if err := providerSrv.DeleteRow(id); err != nil {
		return err
	}

	return c.JSON(200, NewEmptySchema())
}
