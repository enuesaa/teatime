package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/enuesaa/teatime-app/backend/repository"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/enuesaa/teatime-app/backend/binding"
)

func GetAppearance(c *gin.Context) {
	var body v1.SettingGetAppearanceRequest
	if !binding.Validate(c, &body) {
		return
	}

	redis := repository.RedisRepository {}
	value := redis.Get("aaa")
	c.JSON(http.StatusOK, v1.SettingGetAppearanceResponse {
		Message: value,
	})
}

func PutAppearance(c *gin.Context) {
	var body v1.SettingPutAppearanceRequest
	if !binding.Validate(c, &body) {
		return
	}
	redis := repository.RedisRepository {}
	redis.Set("aaa", "bbb")
	c.JSON(http.StatusOK, v1.SettingPutAppearanceResponse {})
}
