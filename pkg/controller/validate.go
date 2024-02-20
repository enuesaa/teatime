package controller

import (
	"fmt"

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

func ValidateMap(c echo.Context, reqbody *map[string]interface{}, rules map[string]interface{}) error {
	if err := c.Bind(&reqbody); err != nil {
		return err
	}

	v := validator.New()
	if errs := v.ValidateMap(*reqbody, rules); len(errs) > 0 {
		return fmt.Errorf("validation error")
	}

	return nil
}
