package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello alworld",
	})
}