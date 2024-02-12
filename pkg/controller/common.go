package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ListResponse[T any] struct {
	Items []T `json:"items"`
}
func NewListResponse[T any]() ListResponse[T] {
	return ListResponse[T] {
		Items: make([]T, 0),
	}
}

type DescribeResponse[T any] struct {
	Data T `json:"data"`
}
func NewDescribeResponse[T any]() DescribeResponse[T] {
	return DescribeResponse[T] {
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
