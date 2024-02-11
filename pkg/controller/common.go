package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ApiResponse[T any] struct {
	Data T `json:"data"`
}
type EmptySchema struct {}
type IdSchema struct {
	Id string `json:"id"`
}

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
