package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/enuesaa/teatime-app/backend/validate"
	"github.com/enuesaa/teatime-app/backend/service"
)


func GetBookmark(c *gin.Context) {
	var body v1.GetBookmarkRequest
	if err := validate.Validate(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := body.Id

	var bookmarkSrv = &service.BookmarkService{}
	bookmarkSrv.Get(c, id)
	
	c.JSON(http.StatusOK, v1.GetBookmarkResponse {
		Id: id,
		Name: "aaa",
		Url: "",
	})
}

func AddBookmark(c *gin.Context) {
	var body v1.AddBookmarkRequest
	if err := validate.Validate(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var bookmarkSrv = &service.BookmarkService{}
	bookmarkSrv.Create(c, service.Bookmark { Name: "aaaaaaa" })
	
	c.JSON(http.StatusOK, v1.AddBookmarkResponse {})
}
