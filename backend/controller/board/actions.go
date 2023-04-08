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
	c.JSON(200, v1.CheckinResponse{ Id: id })
}

func (ctl *BoardController) ListTimeline(c *gin.Context) {
	var body v1.ListTimelineRequest
	if !binding.Validate(c, &body) {
		return
	}
	id := body.Id
	page := body.Page

	checkins := ctl.board().ListCheckins(id)
	items := make([]*v1.ListTimelineResponse_Item, 0)
	for _, v := range checkins {
		items = append(items, &v1.ListTimelineResponse_Item{
			Time: v.Time,
		})
	}
	c.JSON(200, v1.ListTimelineResponse{
		Id: id,
		Page: page,
		Items: items,
	})
}

func (ctl *BoardController) Archive(c *gin.Context) {
	var body v1.ArchiveBoardRequest
	if !binding.Validate(c, &body) {
		return
	}
	id := body.Id

	ctl.board().Archive(id)
	c.JSON(200, v1.ArchiveBoardResponse{ Id: id })
}

func (ctl *BoardController) UnArchive(c *gin.Context) {
	var body v1.UnArchiveBoardRequest
	if !binding.Validate(c, &body) {
		return
	}
	id := body.Id

	ctl.board().UnArchive(id)
	c.JSON(200, v1.UnArchiveBoardResponse{ Id: id })
}
