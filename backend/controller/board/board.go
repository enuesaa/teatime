package board

import (
	"github.com/enuesaa/teatime-app/backend/binding"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/enuesaa/teatime-app/backend/service"
	"github.com/enuesaa/teatime-app/backend/repository"
	"github.com/gin-gonic/gin"
)

type BoardController struct {
	// TODO validator をここに持たせる
	BoardSrv *service.BoardService
}

func (ctl *BoardController) board() *service.BoardService {
	if ctl.BoardSrv == nil {
		ctl.BoardSrv = &service.BoardService{
			RedisRepo: &repository.RedisRepository{},
		}
	}
	return ctl.BoardSrv
}

// func (ctl *BoardController) ListRefactor(c *gin.Context) {
// 	var body v1.ListBoardsRequest
// 	ctl.validator.bind(c, &body) // https://qiita.com/tobita0000/items/d2309cc3f0dd32006ead
// 	ctl.presenter.list(200, 1, data)
// }

func (ctl *BoardController) List(c *gin.Context) {
	var body v1.ListBoardsRequest
	if !binding.Validate(c, &body) {
		return
	}

	list := ctl.board().List()
	items := make([]*v1.ListBoardsResponse_Item, 0)
	for _, v := range list {
		items = append(items, &v1.ListBoardsResponse_Item{
			Id:   v.Id,
			Name: v.Name,
			Description: v.Description,
			Start: v.Start,
			End: v.End,
		})
	}

	c.JSON(200, v1.ListBoardsResponse{
		Page:  1,
		Items: items,
	})
}

func (ctl *BoardController) Get(c *gin.Context) {
	var body v1.GetBoardRequest
	if !binding.Validate(c, &body) {
		return
	}
	id := body.Id

	data := ctl.board().Get(id)
	c.JSON(200, v1.GetBoardResponse{
		Id:   id,
		Name: data.Name,
		Description: data.Description,
		Start: data.Start,
		End: data.End,
	})
}

func (ctl *BoardController) Add(c *gin.Context) {
	var body v1.AddBoardRequest
	if !binding.Validate(c, &body) {
		return
	}

	id := ctl.board().Create(service.Board{
		Name: body.Name,
		Description: body.Description,
		Start: body.Start,
		End: body.End,
	})
	c.JSON(200, v1.AddBoardResponse{Id: id})
}

func (ctl *BoardController) Update(c *gin.Context) {
	var body v1.UpdateBoardRequest
	if !binding.Validate(c, &body) {
		return
	}
	id := body.Id

	ctl.board().Update(id, service.Board{
		Name: body.Name,
		Description: body.Description,
		Start: body.Start,
		End: body.End,
	})
	c.JSON(200, v1.UpdateBoardResponse{Id: id})
}

func (ctl *BoardController) Delete(c *gin.Context) {
	var body v1.DeleteBoardRequest
	if !binding.Validate(c, &body) {
		return
	}
	id := body.Id

	ctl.board().Delete(id)
	c.JSON(200, v1.DeleteBoardResponse{})
}
