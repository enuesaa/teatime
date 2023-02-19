package musics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Spotify(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello alworld",
	})
}
