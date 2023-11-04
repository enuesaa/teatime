package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AbortOnError(c *gin.Context, err error) {
	fmt.Println(err.Error())
	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	c.Abort()
}
