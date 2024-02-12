package controller

import (
	"github.com/labstack/echo/v4"
)

type ProviderSchema struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Command string `json:"command"`
}
func ListProviders(c echo.Context) error {
	res := NewListResponse[ProviderSchema]()
	return c.JSON(200, res)
}

func DescribeProvider(c echo.Context) error {
	res := NewDescribeResponse[ProviderSchema](ProviderSchema{})
	return c.JSON(200, res)
}
