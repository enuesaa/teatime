package main

import (
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

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(jsonMiddleware())

	base := router.Group("/api")
	{
		settingRoute := base.Group("/v1.Setting")
		settingCtl := controller.SettingController {}
		settingRoute.POST("/GetAppearance", settingCtl.Get)
		settingRoute.POST("/PutAppearance", settingCtl.Put)

		bookmarkRoute := base.Group("/v1.Bookmark")
		bookmarkCtl := controller.BookmarkController {}
		bookmarkRoute.POST("/ListBookmarks", bookmarkCtl.List)
		bookmarkRoute.POST("/GetBookmark", bookmarkCtl.Get)
		bookmarkRoute.POST("/AddBookmark", bookmarkCtl.Add)
		
		feedRoute := base.Group("/v1.Feed")
		feedCtl := controller.FeedController {}
		feedRoute.POST("/AddFeed", feedCtl.Add)
	}
	return router
}
