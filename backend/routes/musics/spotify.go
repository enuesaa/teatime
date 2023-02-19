package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SpotifyController struct{}

func (ctrl SpotifyController) One(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello alworld",
	})
}
