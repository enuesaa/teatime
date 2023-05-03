package board

import (
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/gin-gonic/gin"
	"github.com/enuesaa/teatime-app/backend/handle"
)

func (ctl *BoardController) Checkin(c *gin.Context) {
	var req v1.CheckinRequest
	handler := handle.Bind(c, &req)
	handler.Data("Id", func() any {
		id := req.Id
		ctl.board().Checkin(id)
		return id
	})
	handler.Render(&v1.CheckinResponse{})
}

func (ctl *BoardController) ListTimeline(c *gin.Context) {
	var req v1.ListTimelineRequest
	handler := handle.Bind(c, &req)
	handler.Data("Id", func() any {
		return req.Id
	})
	handler.Data("Page", func() any {
		return req.Page
	})
	handler.Data("Items", func() any {
		id := req.Id
		return ctl.board().ListCheckins(id)
	})
	handler.Render(&v1.ListTimelineResponse{})
}

func (ctl *BoardController) Archive(c *gin.Context) {
	var req v1.ArchiveBoardRequest
	handler := handle.Bind(c, &req)
	handler.Data("Id", func() any {
		id := req.Id
		ctl.board().Archive(id)
		return id
	})
	handler.Render(&v1.ArchiveBoardResponse{})
}

func (ctl *BoardController) UnArchive(c *gin.Context) {
	var req v1.UnArchiveBoardRequest
	handler := handle.Bind(c, &req)
	handler.Data("Id", func() any {
		id := req.Id
		ctl.board().UnArchive(id)
		return id
	})
	handler.Render(&v1.UnArchiveBoardResponse{})
}
