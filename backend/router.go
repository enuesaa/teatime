package main

import (
	"github.com/enuesaa/teatime-app/backend/controller/setting"
	"github.com/enuesaa/teatime-app/backend/controller/bookmark"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(jsonMiddleware())
	base := router.Group("/api")
	{
		settingRoute := base.Group("/v1.Setting")
		settingRoute.POST("/GetAppearance", setting.GetAppearance)
		settingRoute.POST("/PutAppearance", setting.PutAppearance)

		bookmarkRoute := base.Group("/v1.Bookmark")
		bookmarkRoute.POST("/AddBookmark", bookmark.AddBookmark)
	}
	return router
}