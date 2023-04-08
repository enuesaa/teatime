package board

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

func (ctl *BoardController) List (c *gin.Context) {
	var body v1.ListBoardsRequest
	if !binding.Validate(c, &body) { return }

	list := ctl.board().List()
	items := make([]*v1.ListBoardsResponse_Item, 0)
	for _, v := range list {
		items = append(items, &v1.ListBoardsResponse_Item {
			Id: v.Id,
			Name: v.Name,
		})
	}

	c.JSON(200, v1.ListBoardsResponse {
		Page: 1,
		Items: items,
	})
}

func (ctl *BoardController) Get (c *gin.Context) {
	var body v1.GetBoardRequest
	if !binding.Validate(c, &body) { return }
	id := body.Id

	data := ctl.board().Get(id)	
	c.JSON(200, v1.GetBoardResponse {
		Id: id,
		Name: data.Name,
	})
}

func (ctl *BoardController) Add (c *gin.Context) {
	var body v1.AddBoardRequest
	if !binding.Validate(c, &body) { return }

	id := ctl.board().Create(service.Board {
		Name: body.Name,
	})
	c.JSON(200, v1.AddBoardResponse { Id: id })
}

func (ctl *BoardController) Update (c *gin.Context) {
	var body v1.UpdateBoardRequest
	if !binding.Validate(c, &body) { return }
	id := body.Id

	ctl.board().Update(id, service.Board {
		Name: body.Name,
	})
	c.JSON(200, v1.UpdateBoardResponse { Id: id })
}

func (ctl *BoardController) Delete (c *gin.Context) {
	var body v1.DeleteBoardRequest
	if !binding.Validate(c, &body) { return }
	id := body.Id

	ctl.board().Delete(id)
	c.JSON(200, v1.DeleteBoardResponse {})
}


func (ctl *BoardController) Checkin (c *gin.Context) {
	var body v1.CheckinRequest
	if !binding.Validate(c, &body) { return }
	// create item
	c.JSON(200, v1.CheckinResponse {})
}

func (ctl *BoardController) ListTimeline (c *gin.Context) {
	var body v1.ListTimelineRequest
	if !binding.Validate(c, &body) { return }
	// list items
	c.JSON(200, v1.ListTimelineResponse {})
}

func (ctl *BoardController) Archive (c *gin.Context) {
	var body v1.ArchiveBoardRequest
	if !binding.Validate(c, &body) { return }
	id := body.Id

	ctl.board().Archive(id)
	c.JSON(200, v1.ArchiveBoardResponse {})
}

func (ctl *BoardController) UnArchive (c *gin.Context) {
	var body v1.UnArchiveBoardRequest
	if !binding.Validate(c, &body) { return }
	id := body.Id

	ctl.board().UnArchive(id)
	c.JSON(200, v1.UnArchiveBoardResponse {})
}
