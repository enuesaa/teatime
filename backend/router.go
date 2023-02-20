package main

import (
	"github.com/enuesaa/teatime-app/backend/controllers/settings"
	"github.com/enuesaa/teatime-app/backend/controllers/lounge"
	"github.com/enuesaa/teatime-app/backend/controllers/bookshelf"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	base := router.Group("/api")
	{
		settingsRoute := base.Group("/settings")
		settingsRoute.GET("/appearance", settings.GetAppearance)
		settingsRoute.PUT("/appearance", settings.PutAppearance)
		
		loungeRoute := base.Group("/lounge")
		loungeRoute.GET("/spotify", lounge.CallSpotifyApi)

		bookshelfRoute := base.Group("/bookshelf")
		bookshelfRoute.GET("/list", bookshelf.ListBooks)
	}

	return router
}
