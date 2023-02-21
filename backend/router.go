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
		settingsRoute := base.Group("/v1.SettingsService")
		settingsRoute.POST("/GetAppearance", settings.GetAppearance)
		
		loungeRoute := base.Group("/v1.LoungeService")
		loungeRoute.POST("/Spotify", lounge.CallSpotifyApi)

		bookshelfRoute := base.Group("/v1.BookShelfService")
		bookshelfRoute.POST("/List", bookshelf.ListBooks)
	}

	return router
}
