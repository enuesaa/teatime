package setting

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/enuesaa/teatime-app/backend/repository"
)

func GetAppearance(ctx *gin.Context) {
	value := repository.GetValue(ctx, "aaa")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello alworld",
		"redisvalue": value,
	})
}

func PutAppearance(ctx *gin.Context) {
	repository.SetValue(ctx, "aaa", "bbb")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello alworld",
	})
}
