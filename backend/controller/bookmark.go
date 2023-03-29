package controller

import (
	"fmt"
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

	var bookmarkSrv = service.NewBookmarkService()
	data := bookmarkSrv.Get(id)
	fmt.Println("%v", data)
	
	c.JSON(http.StatusOK, v1.GetBookmarkResponse {
		Id: id,
		Name: data.Name,
		Url: "",
	})
}

func AddBookmark(c *gin.Context) {
	var body v1.AddBookmarkRequest
	if err := validate.Validate(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var bookmarkSrv = service.NewBookmarkService()
	id := bookmarkSrv.Create(service.Bookmark { Name: "aaaaaaa" })
	
	c.JSON(http.StatusOK, v1.AddBookmarkResponse {
		Id: id,
	})
}
