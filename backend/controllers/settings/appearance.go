package settings

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAppearance(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello alworld",
	})
}

func PutAppearance(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello alworld",
	})
}
