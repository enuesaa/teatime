package controller

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Validate(c echo.Context, reqbody interface{}) error {
	if err := c.Bind(&reqbody); err != nil {
		return err
	}

	v := validator.New()
	if err := v.Struct(reqbody); err != nil {
		return err
	}

	return nil
}

type ValidationErrors struct {
	error
	Errors []ValidationError `json:"errors"`
}
type ValidationError struct {
	Path string `json:"path"`
	Message string `json:"message"`
}

func ValidateMap(c echo.Context, reqbody *plug.Values, rules map[string]interface{}) error {
	if err := c.Bind(&reqbody); err != nil {
		return err
	}

	validateData := make(map[string]interface{}, 0)
	for name, val := range *reqbody {
		validateData[name] = val
	}
	v := validator.New()
	if errs := v.ValidateMap(validateData, rules); len(errs) > 0 {
		vaerrs := ValidationErrors{
			Errors: make([]ValidationError, 0),
		}
		for path, errdetail := range errs {
			err := errdetail.(validator.ValidationErrors)[0]
			vaerrs.Errors = append(vaerrs.Errors, ValidationError{
				Path: path,
				Message: fmt.Sprintf("something wrong with %s", err.Tag()),
			})
		}
		return vaerrs
	}

	return nil
}
