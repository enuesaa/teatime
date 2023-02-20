package lounge

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CallSpotifyApi(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello alworld",
	})
}
