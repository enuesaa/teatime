package controller

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type AppInternalServerErr struct {
	error
	Errors []AppErrItem `json:"errors"`
}
type AppErrItem struct {
	Path string `json:"path"`
	Message string `json:"message"`
}
func HandleError(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}
		var validationErrs AppValidationError
		if errors.As(err, &validationErrs) {
			return c.JSON(422, err)
		}
		appErr := AppInternalServerErr {
			Errors: make([]AppErrItem, 0),
		}
		appErr.Errors = append(appErr.Errors, AppErrItem {
			Path: "$",
			Message: "Internal Server Error",
		})
		return c.JSON(500, appErr)
	}
}
