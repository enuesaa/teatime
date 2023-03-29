package main

import (
	"fmt"
	"github.com/enuesaa/teatime-app/backend/controller"
	"github.com/gin-gonic/gin"
)

func jsonMiddleware() gin.HandlerFunc {
	// https://stackoverflow.com/questions/41109065/golang-gin-gonic-content-type-not-setting-to-application-json-with-c-json
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func errorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errs := c.Errors.Errors()
		if len(errs) > 0 {
			c.JSON(400, gin.H{"error": errs[0] })
		}
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(jsonMiddleware())
	router.Use(errorMiddleware())

	base := router.Group("/api")
	{
		settingRoute := base.Group("/v1.Setting")
		settingRoute.POST("/GetAppearance", controller.GetAppearance)
		settingRoute.POST("/PutAppearance", controller.PutAppearance)

		bookmarkRoute := base.Group("/v1.Bookmark")
		bookmarkRoute.POST("/ListBookmarks", controller.ListBookmarks)
		bookmarkRoute.POST("/GetBookmark", controller.GetBookmark)
		bookmarkRoute.POST("/AddBookmark", controller.AddBookmark)
		
		feedRoute := base.Group("/v1.Feed")
		feedRoute.POST("/AddFeed", controller.AddFeed)
	}
	return router
}
