package controller

import (
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
