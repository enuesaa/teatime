package board

import (
	"github.com/enuesaa/teatime-app/backend/binding"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/gin-gonic/gin"
)

func (ctl *BoardController) Checkin(c *gin.Context) {
	var body v1.CheckinRequest
	if !binding.Validate(c, &body) {
		return
	}
	id := body.Id

	ctl.board().Checkin(id)
	c.JSON(200, v1.CheckinResponse{})
}

func (ctl *BoardController) ListTimeline(c *gin.Context) {
	var body v1.ListTimelineRequest
	if !binding.Validate(c, &body) {
		return
	}
	c.JSON(200, v1.ListTimelineResponse{})
}

func (ctl *BoardController) Archive(c *gin.Context) {
	var body v1.ArchiveBoardRequest
	if !binding.Validate(c, &body) {
		return
	}
	id := body.Id

	ctl.board().Archive(id)
	c.JSON(200, v1.ArchiveBoardResponse{})
}

func (ctl *BoardController) UnArchive(c *gin.Context) {
	var body v1.UnArchiveBoardRequest
	if !binding.Validate(c, &body) {
		return
	}
	id := body.Id

	ctl.board().UnArchive(id)
	c.JSON(200, v1.UnArchiveBoardResponse{})
}
