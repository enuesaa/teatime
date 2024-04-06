package controller

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AppValidationError struct {
	error
	Errors []AppErrItem `json:"errors"`
}

func Validate(c echo.Context, reqbody *plug.Value, rules map[string]interface{}) error {
	if err := c.Bind(&reqbody); err != nil {
		return err
	}

	// this is for validation
	data := make(map[string]interface{}, 0)
	for name, val := range *reqbody {
		data[name] = val
	}
	errs := validator.New().ValidateMap(data, rules)
	if len(errs) == 0 {
		return nil
	}

	apperr := NewAppErr()
	for path, err := range errs {
		err := err.(validator.ValidationErrors)[0]
		apperr.Errors = append(apperr.Errors, AppErrItem{
			Path:    path,
			Message: fmt.Sprintf("something wrong with %s", err.Tag()),
		})
	}

	return apperr
}
