package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/enuesaa/teatime-app/backend/binding"
	"github.com/enuesaa/teatime-app/backend/service"
)


type FeedController struct {
	FeedSrv *service.FeedService
}

func (ctl *FeedController) feed () *service.FeedService {
	if ctl.FeedSrv == nil {
		ctl.FeedSrv = service.NewFeedService()
	}
	return ctl.FeedSrv
}

func (ctl *FeedController) Add (c *gin.Context) {
	var body v1.AddFeedRequest
	if !binding.Validate(c, &body) { return }

	ctl.feed().Fetch(body.Url)
	
	c.JSON(200, v1.AddFeedResponse {})
}
