package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/enuesaa/teatime-app/backend/validate"
	"github.com/enuesaa/teatime-app/backend/service/feed"
)

func AddFeed(c *gin.Context) {
	var body v1.AddFeedRequest
	if err := validate.Validate(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	feed.FetchFeed(body.Url)
	
	c.JSON(http.StatusOK, v1.AddFeedResponse {})
}
