package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ApiResponse[T any] struct {
	Items []T `json:"items"`
	Data T `json:"data"`
}

func NewListResponse[T any]() ApiResponse[T] {
	return ApiResponse[T] {
		Items: make([]T, 0),
	}
}
func NewDescribeResponse[T any]() ApiResponse[T] {
	return ApiResponse[T] {
		Data: *new(T),
	}
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
