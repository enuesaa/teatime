package bookmark

import (
	"github.com/enuesaa/teatime-app/backend/binding"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/enuesaa/teatime-app/backend/service"
	"github.com/enuesaa/teatime-app/backend/repository"
	"github.com/gin-gonic/gin"
)

type BookmarkController struct {
	BookmarkSrv *service.BookmarkService
}

func (ctl *BookmarkController) bookmark() *service.BookmarkService {
	if ctl.BookmarkSrv == nil {
		ctl.BookmarkSrv = &service.BookmarkService{
			RedisRepo: &repository.RedisRepository{},
		}
	}
	return ctl.BookmarkSrv
}

func (ctl *BookmarkController) List(c *gin.Context) {
	var body v1.ListBookmarksRequest
	if !binding.Validate(c, &body) {
		return
	}

	list := ctl.bookmark().List()
	items := make([]*v1.ListBookmarksResponse_Item, 0)
	for _, v := range list {
		items = append(items, &v1.ListBookmarksResponse_Item{
			Id:   v.Id,
			Name: v.Name,
			Url:  v.Url,
		})
	}

	c.JSON(200, v1.ListBookmarksResponse{
		Page:  1,
		Items: items,
	})
}

func (ctl *BookmarkController) Get(c *gin.Context) {
	var body v1.GetBookmarkRequest
	if !binding.Validate(c, &body) {
		return
	}
	id := body.Id

	data := ctl.bookmark().Get(id)
	c.JSON(200, v1.GetBookmarkResponse{
		Id:   id,
		Name: data.Name,
		Url:  data.Url,
	})
}

func (ctl *BookmarkController) Add(c *gin.Context) {
	var body v1.AddBookmarkRequest
	if !binding.Validate(c, &body) {
		return
	}

	id := ctl.bookmark().Create(service.Bookmark{
		Name: body.Name,
		Url:  body.Url,
	})
	c.JSON(200, v1.AddBookmarkResponse{Id: id})
}

func (ctl *BookmarkController) Update(c *gin.Context) {
	var body v1.UpdateBookmarkRequest
	if !binding.Validate(c, &body) {
		return
	}
	id := body.Id

	ctl.bookmark().Update(id, service.Bookmark{
		Name: body.Name,
		Url:  body.Url,
	})
	c.JSON(200, v1.UpdateBookmarkResponse{Id: id})
}

func (ctl *BookmarkController) Delete(c *gin.Context) {
	var body v1.DeleteBookmarkRequest
	if !binding.Validate(c, &body) {
		return
	}
	id := body.Id

	ctl.bookmark().Delete(id)
	c.JSON(200, v1.DeleteBookmarkResponse{})
}
