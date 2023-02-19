package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MusicsController struct{}

func (ctrl MusicsController) One(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello alworld",
	})
}
