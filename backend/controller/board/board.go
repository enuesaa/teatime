package board

import (
	"fmt"

	"github.com/enuesaa/teatime-app/backend/binding"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/enuesaa/teatime-app/backend/repository"
	"github.com/enuesaa/teatime-app/backend/service"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type BoardController struct {
	BoardSrv *service.BoardService
	data map[string]interface{}
}

func (ctl *BoardController) board() *service.BoardService {
	if ctl.BoardSrv == nil {
		ctl.BoardSrv = &service.BoardService{
			RedisRepo: &repository.RedisRepository{},
		}
	}
	return ctl.BoardSrv
}

func (ctl *BoardController) Bind(c *gin.Context, req binding.WithValidator) {
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
	}
	ctl.data = make(map[string]interface{}, 0)
}

type HandleFunc func() any
func (ctl *BoardController) Handle(c *gin.Context, name string, fn HandleFunc) {
	if (c.IsAborted()) {
		return;
	}
	ctl.data[name] = fn()
}

func (ctl *BoardController) Render(c *gin.Context, response any) {
	if (c.IsAborted()) {
		return;
	}
	if val, ok := ctl.data["*"]; ok {
		mapstructure.Decode(val, response)
	} else {
		mapstructure.Decode(ctl.data, response)
	}
	c.JSON(200, response)
}


func (ctl *BoardController) List(c *gin.Context) {
	ctl.Bind(c, &v1.ListBoardsRequest{})
	ctl.Handle(c, "Page", func() any {
		return 1
	})
	ctl.Handle(c, "Items", func() any {
		res := ctl.board().List()
		for _, v := range res {
			fmt.Printf("{%+v}", v)
		}
		return res
	})
	ctl.Render(c, &v1.ListBoardsResponse{})
}

func (ctl *BoardController) Get(c *gin.Context) {
	var req v1.GetBoardRequest
	ctl.Bind(c, &req)
	ctl.Handle(c, "*", func() any {
		id := req.Id
		return ctl.board().Get(id)
	})
	ctl.Render(c, &v1.GetBoardResponse{})
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
