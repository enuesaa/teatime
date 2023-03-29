package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/enuesaa/teatime-app/backend/binding"
	"github.com/enuesaa/teatime-app/backend/service"
)

func ListBookmarks(c *gin.Context) {
	var body v1.ListBookmarksRequest
	if !binding.Validate(c, &body) {
		return
	}

	var bookmarkSrv = service.NewBookmarkService()
	list := bookmarkSrv.List()
	items := make([]*v1.ListBookmarksResponse_Item, 0)
	for _, v := range list {
		items = append(items, &v1.ListBookmarksResponse_Item {
			Id: v.Id,
			Name: v.Name,
			Url: v.Url,
		})
	}

	c.JSON(200, v1.ListBookmarksResponse {
		Page: 1,
		Items: items,
	})
}

func GetBookmark(c *gin.Context) {
	var body v1.GetBookmarkRequest
	if !binding.Validate(c, &body) {
		return
	}
	id := body.Id

	var bookmarkSrv = service.NewBookmarkService()
	data := bookmarkSrv.Get(id)	
	c.JSON(200, v1.GetBookmarkResponse {
		Id: id,
		Name: data.Name,
		Url: "",
	})
}

func AddBookmark(c *gin.Context) {
	var body v1.AddBookmarkRequest
	if !binding.Validate(c, &body) {
		return
	}

	var bookmarkSrv = service.NewBookmarkService()
	id := bookmarkSrv.Create(service.Bookmark {
		Name: body.Name,
		Url: body.Url,
	})

	c.JSON(200, v1.AddBookmarkResponse { Id: id })
}
