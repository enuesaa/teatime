package controller

import (
	"github.com/labstack/echo/v4"
)

type ProviderRowSchema struct {
	Id string `json:"id"` // rowid
	Columns map[string]string `json:"columns"` // TODO boolean を表現できるようにする
}

func ListProviderRows(c echo.Context) error {
	res := NewListResponse[ProviderRowSchema]()
	return c.JSON(200, res)
}

func DescribeProviderRow(c echo.Context) error {
	res := NewDescribeResponse[ProviderRowSchema]()
	return c.JSON(200, res)
}

func CreateProviderRow(c echo.Context) error {
	return c.JSON(200, NewIdSchema("a"))
}

func UpdateProviderRow(c echo.Context) error {
	return c.JSON(200, NewIdSchema("a"))
}

func DeleteProviderRow(c echo.Context) error {
	return c.JSON(200, NewEmptySchema())
}
