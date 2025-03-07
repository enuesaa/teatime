package middleware

import (
	"github.com/labstack/echo/v4"
)

type ApiResponse struct {
	Items interface{} `json:"items"`
}

func HandleData(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			return err
		}

		// check that reponse body is already set.
		if c.Response().Committed {
			return nil
		}

		items := c.Get("items")
		if items != nil {
			res := ApiResponse{
				Items: items,
			}
			return c.JSON(200, res)
		}

		data := c.Get("data")
		return c.JSON(200, data)
	}
}
