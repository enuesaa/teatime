package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ApiResponse[T any] struct {
	Data T `json:"data"`
}
type EmptySchema struct {}

func AbortOnError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	c.Abort()
}

func Validate(c *gin.Context, reqbody interface{}) error {
	if err := c.ShouldBindJSON(&reqbody); err != nil {
		return err
	}

	v := validator.New()
	if err := v.Struct(reqbody); err != nil {
		return err
	}

	return nil
}
