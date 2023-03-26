package validate

import (
	"github.com/gin-gonic/gin"
)

type WithValidator interface {
	Validate() error
}

// see https://github.com/eddycjy/go-gin-example/blob/master/pkg/app/form.go
func Validate(c *gin.Context, body WithValidator) error {
	if err := c.ShouldBindJSON(&body); err != nil {
		return err
	}
	if err := body.Validate(); err != nil {
		return err
	}
	return nil
}