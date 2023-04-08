package binding

import (
	"github.com/gin-gonic/gin"
)

type WithValidator interface {
	Validate() error
}

// see https://github.com/eddycjy/go-gin-example/blob/master/pkg/app/form.go
func Validate(c *gin.Context, body WithValidator) bool {
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return false
	}
	if err := body.Validate(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return false
	}
	return true
}
