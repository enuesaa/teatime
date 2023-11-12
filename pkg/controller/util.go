package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AbortOnError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	c.Abort()
}
