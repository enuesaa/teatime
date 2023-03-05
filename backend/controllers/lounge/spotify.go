package lounge

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/enuesaa/teatime-app/backend/services/musics"
)

func CallSpotifyApi(ctx *gin.Context) {
	data := musics.CallSpotifyApi(ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
	})
}
