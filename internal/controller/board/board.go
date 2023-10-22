package board

// import (
// 	"github.com/enuesaa/teatime/internal/handle"
// 	"github.com/enuesaa/teatime/internal/repository"
// 	"github.com/enuesaa/teatime/internal/service"
// 	"github.com/gin-gonic/gin"
// )

// type BoardController struct {
// 	BoardSrv *service.BoardService
// 	data map[string]interface{}
// }

// func (ctl *BoardController) board() *service.BoardService {
// 	if ctl.BoardSrv == nil {
// 		ctl.BoardSrv = &service.BoardService{
// 			RedisRepo: &repository.RedisRepository{},
// 		}
// 	}
// 	return ctl.BoardSrv
// }

// func (ctl *BoardController) List(c *gin.Context) {
// 	handler := handle.Bind(c, &v1.ListBoardsRequest{})
// 	handler.Data("Page", func() any {
// 		return 1
// 	})
// 	handler.Data("Items", func() any {
// 		return ctl.board().List()
// 	})
// 	handler.Render(&v1.ListBoardsResponse{})
// }

// func (ctl *BoardController) Get(c *gin.Context) {
// 	var req v1.GetBoardRequest
// 	handler := handle.Bind(c, &req)
// 	handler.Data("*", func() any {
// 		id := req.Id
// 		return ctl.board().Get(id)
// 	})
// 	handler.Render(&v1.GetBoardResponse{})
// }

// func (ctl *BoardController) Add(c *gin.Context) {
// 	var req v1.AddBoardRequest
// 	handler := handle.Bind(c, &req)
// 	handler.Data("Id", func() any {
// 		id := ctl.board().Create(service.Board{
// 			Name: req.Name,
// 			Description: req.Description,
// 			Start: req.Start,
// 			End: req.End,
// 		})
// 		return id
// 	})
// 	handler.Render(&v1.AddBoardResponse{})
// }

// func (ctl *BoardController) Update(c *gin.Context) {
// 	var req v1.UpdateBoardRequest
// 	handler := handle.Bind(c, &req)
// 	handler.Data("Id", func() any {
// 		id := req.Id
// 		ctl.board().Update(id, service.Board{
// 			Name: req.Name,
// 			Description: req.Description,
// 			Start: req.Start,
// 			End: req.End,
// 		})
// 		return id
// 	})
// 	handler.Render(&v1.UpdateBoardResponse{})
// }

// func (ctl *BoardController) Delete(c *gin.Context) {
// 	var req v1.DeleteBoardRequest
// 	handler := handle.Bind(c, &req)
// 	handler.Process(func() {
// 		id := req.Id
// 		ctl.board().Delete(id)
// 	})
// 	handler.Render(&v1.DeleteBoardResponse{})
// }
