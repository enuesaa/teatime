package bookmark

// import (
// 	"github.com/enuesaa/teatime/internal/service"
// 	"github.com/enuesaa/teatime/internal/repository"
// 	"github.com/gin-gonic/gin"
// 	"github.com/enuesaa/teatime/internal/handle"
// )

// type BookmarkController struct {
// 	BookmarkSrv *service.BookmarkService
// }

// func (ctl *BookmarkController) bookmark() *service.BookmarkService {
// 	if ctl.BookmarkSrv == nil {
// 		ctl.BookmarkSrv = &service.BookmarkService{
// 			RedisRepo: &repository.RedisRepository{},
// 		}
// 	}
// 	return ctl.BookmarkSrv
// }

// // 一個一個ハンドラーを分ける
// func (ctl *BookmarkController) List(c *gin.Context) {
// 	var req v1.ListBookmarksRequest
// 	handler := handle.Bind(c, &req)
// 	handler.Data("Page", func() any {
// 		return 1
// 	})
// 	handler.Data("Items", func() any {
// 		return ctl.bookmark().List()
// 	})
// 	handler.Render(&v1.ListBookmarksResponse{})
// }

// func (ctl *BookmarkController) Get(c *gin.Context) {
// 	var req v1.GetBookmarkRequest
// 	handler := handle.Bind(c, &req)
// 	handler.Data("*", func() any {
// 		id := req.Id
// 		return ctl.bookmark().Get(id)
// 	})
// 	handler.Render(&v1.GetBookmarkResponse{})
// }

// func (ctl *BookmarkController) Add(c *gin.Context) {
// 	var req v1.AddBookmarkRequest
// 	handler := handle.Bind(c, &req)
// 	handler.Data("Id", func() any {
// 		id := ctl.bookmark().Create(service.Bookmark{
// 			Name: req.Name,
// 			Url:  req.Url,
// 		})
// 		return id
// 	})
// 	handler.Render(&v1.AddBookmarkResponse{})
// }

// func (ctl *BookmarkController) Update(c *gin.Context) {
// 	var req v1.UpdateBookmarkRequest
// 	handler := handle.Bind(c, &req)
// 	handler.Data("Id", func() any {
// 		id := req.Id
// 		ctl.bookmark().Update(id, service.Bookmark{
// 			Name: req.Name,
// 			Url:  req.Url,
// 		})
// 		return id
// 	})
// 	handler.Render(&v1.UpdateBookmarkResponse{})
// }

// func (ctl *BookmarkController) Delete(c *gin.Context) {
// 	var req v1.DeleteBookmarkRequest
// 	handler := handle.Bind(c, &req)
// 	handler.Process(func() {
// 		id := req.Id
// 		ctl.bookmark().Delete(id)
// 	})
// 	handler.Render(&v1.DeleteBookmarkResponse{})
// }
