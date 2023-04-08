package setting

import (
	"net/http"

	"github.com/enuesaa/teatime-app/backend/binding"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/enuesaa/teatime-app/backend/repository"
	"github.com/gin-gonic/gin"
)

type SettingController struct {
	RedisRepo *repository.RedisRepository
}

func (ctl *SettingController) redis() *repository.RedisRepository {
	if ctl.RedisRepo == nil {
		ctl.RedisRepo = &repository.RedisRepository{}
	}
	return ctl.RedisRepo
}

func (ctl *SettingController) Get(c *gin.Context) {
	var body v1.SettingGetAppearanceRequest
	if !binding.Validate(c, &body) {
		return
	}

	value := ctl.redis().Get("aaa")
	c.JSON(http.StatusOK, v1.SettingGetAppearanceResponse{
		Message: value,
	})
}

func (ctl *SettingController) Put(c *gin.Context) {
	var body v1.SettingPutAppearanceRequest
	if !binding.Validate(c, &body) {
		return
	}

	ctl.redis().Set("aaa", "bbb")
	c.JSON(http.StatusOK, v1.SettingPutAppearanceResponse{})
}
