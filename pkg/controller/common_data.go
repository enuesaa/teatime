package controller

import (
	"github.com/labstack/echo/v4"
)

type AppContext struct {
    echo.Context
}
func (c *AppContext) Data(data interface{}) error {
	c.Set("data", data)
	return nil
}

func HandleData(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(&AppContext{c})
		if err != nil {
			return err
		}
		return nil
	}
}
