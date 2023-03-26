package setting

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/enuesaa/teatime-app/backend/repository/redis"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
)

func validateGetAppearance(body *v1.SettingGetAppearanceRequest) bool {
	return true
}
func GetAppearance(ctx *gin.Context) {
	var body v1.SettingGetAppearanceRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if !validateGetAppearance(&body) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	value := redis.GetValue(ctx, "aaa")
	ctx.JSON(http.StatusOK, v1.SettingGetAppearanceResponse {
		Message: value,
	})
}
