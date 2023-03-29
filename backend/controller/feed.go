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

func (ctl *FeedController) Get (c *gin.Context) {
	var body v1.GetFeedRequest
	if !binding.Validate(c, &body) { return }

	c.JSON(200, v1.GetFeedResponse {})
}

func (ctl *FeedController) ListItems (c *gin.Context) {
	var body v1.ListItemsRequest
	if !binding.Validate(c, &body) { return }

	c.JSON(200, v1.ListItemsResponse {})
}

func (ctl *FeedController) GetAppearance (c *gin.Context) {
	var body v1.GetAppearanceRequest
	if !binding.Validate(c, &body) { return }

	c.JSON(200, v1.GetAppearanceResponse {})
}

func (ctl *FeedController) UpdateAppearance (c *gin.Context) {
	var body v1.UpdateAppearanceRequest
	if !binding.Validate(c, &body) { return }

	c.JSON(200, v1.UpdateAppearanceResponse {})
}

func (ctl *FeedController) Fetch (c *gin.Context) {
	var body v1.FetchRequest
	if !binding.Validate(c, &body) { return }

	c.JSON(200, v1.FetchResponse {})
}

func (ctl *FeedController) Delete (c *gin.Context) {
	var body v1.DeleteFeedRequest
	if !binding.Validate(c, &body) { return }

	c.JSON(200, v1.DeleteFeedResponse {})
}
