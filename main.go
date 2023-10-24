package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.Use(func(c *gin.Context) {
		// https://stackoverflow.com/questions/41109065/golang-gin-gonic-content-type-not-setting-to-application-json-with-c-json
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	})

	// bookmarkRoute := base.Group("/v1.Bookmark")
	// bookmarkCtl := bookmark.BookmarkController{}
	// bookmarkRoute.POST("/ListBookmarks", bookmarkCtl.List)
	// bookmarkRoute.POST("/GetBookmark", bookmarkCtl.Get)
	// bookmarkRoute.POST("/AddBookmark", bookmarkCtl.Add)
	// bookmarkRoute.POST("/UpdateBookmark", bookmarkCtl.Update)
	// bookmarkRoute.POST("/DeleteBookmark", bookmarkCtl.Delete)

	app.Run(":3000")
}
