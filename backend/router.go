package main

import (
	"github.com/enuesaa/teatime-app/backend/controller"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(jsonMiddleware())
	base := router.Group("/api")
	{
		settingRoute := base.Group("/v1.Setting")
		settingRoute.POST("/GetAppearance", controller.GetAppearance)
		settingRoute.POST("/PutAppearance", controller.PutAppearance)

		bookmarkRoute := base.Group("/v1.Bookmark")
		bookmarkRoute.POST("/GetBookmark", controller.GetBookmark)
		bookmarkRoute.POST("/AddBookmark", controller.AddBookmark)
	}
	return router
}
