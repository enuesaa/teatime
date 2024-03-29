package controller

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type AppErr struct {
	error
	Errors []AppErrItem `json:"errors"`
}
type AppErrItem struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}
func NewAppErr() AppErr {
	return AppErr{
		Errors: make([]AppErrItem, 0),
	}
}

func HandleError(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}
		var apperr AppErr
		if errors.As(err, &apperr) {
			return c.JSON(422, err)
		}

		apperr = NewAppErr()
		apperr.Errors = append(apperr.Errors, AppErrItem{
			Path:    "$",
			Message: "Internal Server Error",
		})
		return c.JSON(500, apperr)
	}
}
