package middleware

import "github.com/labstack/echo/v4"

type Err struct {
	Error string `json:"error"`
}

func HandleError(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}
		// like 404
		if _, ok := err.(*echo.HTTPError); ok {
			return err
		}

		res := Err{
			Error: err.Error(),
		}
		return c.JSON(400, res)
	}
}
