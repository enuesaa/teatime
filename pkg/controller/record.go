package controller

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

func RegisterRecord(c echo.Context) error {
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)

	model := c.Param("model")
	recordName := c.Param("recordName")

	err := providerSrv.Register(model, recordName)
	if err != nil {
		return err
	}

	res := ApiResponse[struct{}] {
		Data: struct{}{},
	}
	return c.JSON(200, res)
}

func GetRecord(c echo.Context) error {
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)

	model := c.Param("model")
	recordName := c.Param("recordName")

	record, err := providerSrv.Get(model, recordName)
	if err != nil {
		return err
	}

	res := ApiResponse[plug.Record] {
		Data: record,
	}
	return c.JSON(200, res)
}

func SetRecord(c echo.Context) error {
	var record plug.Record
	if err := c.Bind(&record); err != nil {
		return err
	}
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)

	model := c.Param("model")
	recordName := c.Param("recordName")

	err := providerSrv.Set(model, recordName, record)
	if err != nil {
		return err
	}

	res := ApiResponse[struct{}] {
		Data: struct{}{},
	}
	return c.JSON(200, res)
}

func DelRecord(c echo.Context) error {
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)

	model := c.Param("model")
	recordName := c.Param("recordName")

	err := providerSrv.Del(model, recordName)
	if err != nil {
		return err
	}

	res := ApiResponse[struct{}] {
		Data: struct{}{},
	}
	return c.JSON(200, res)
}
