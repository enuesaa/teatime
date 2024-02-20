package controller

import (
	"github.com/labstack/echo/v4"
)

func WithData(c echo.Context, data interface{}) error {
	c.Set("data", data)
	return nil
}
func HandleData(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			return err
		}
		data := c.Get("data")
		if data == nil {
			data = struct{}{}
		}
		res := ApiResponse {
			Data: data,
		}
		return c.JSON(200, res)
	}
}
