package controller

import (
	"github.com/labstack/echo/v4"
)

type ProviderConfigSchema struct {
	Schema interface{} `json:"schema"` // TODO
}
func DescribeProviderConfig(c echo.Context) error {
	res := NewDescribeResponse[ProviderConfigSchema]()
	return c.JSON(200, res)
}
