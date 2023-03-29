package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/enuesaa/teatime-app/backend/binding"
	"github.com/enuesaa/teatime-app/backend/service"
)

func AddFeed(c *gin.Context) {
	var body v1.AddFeedRequest
	if !binding.Validate(c, &body) {
		return
	}
	feedSrv := service.NewFeedService()
	feedSrv.Fetch(body.Url)
	
	c.JSON(200, v1.AddFeedResponse {})
}
