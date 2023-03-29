package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/enuesaa/teatime-app/backend/binding"
	"github.com/enuesaa/teatime-app/backend/service"
)

type BoardController struct {
	BoardSrv *service.BoardService
}

func (ctl *BoardController) board () *service.BoardService {
	if ctl.BoardSrv == nil {
		ctl.BoardSrv = service.NewBoardService()
	}
	return ctl.BoardSrv
}

func (ctl *BoardController) Add (c *gin.Context) {
	var body v1.AddBoardRequest
	if !binding.Validate(c, &body) { return }
	c.JSON(200, v1.AddBoardResponse {})
}

func (ctl *BoardController) Checkin (c *gin.Context) {
	var body v1.CheckinRequest
	if !binding.Validate(c, &body) { return }
	c.JSON(200, v1.CheckinResponse {})
}

func (ctl *BoardController) ListTimeline (c *gin.Context) {
	var body v1.ListTimelineRequest
	if !binding.Validate(c, &body) { return }
	c.JSON(200, v1.ListTimelineResponse {})
}

func (ctl *BoardController) Archive (c *gin.Context) {
	var body v1.ArchiveBoardRequest
	if !binding.Validate(c, &body) { return }
	c.JSON(200, v1.ArchiveBoardResponse {})
}

func (ctl *BoardController) UnArchive (c *gin.Context) {
	var body v1.UnArchiveBoardRequest
	if !binding.Validate(c, &body) { return }
	c.JSON(200, v1.UnArchiveBoardResponse {})
}
