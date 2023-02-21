package main

import (
	"github.com/enuesaa/teatime-app/backend/controllers/settings"
	"github.com/enuesaa/teatime-app/backend/controllers/lounge"
	"github.com/enuesaa/teatime-app/backend/controllers/bookshelf"
	"github.com/gin-gonic/gin"
)

// https://stackoverflow.com/questions/41109065/golang-gin-gonic-content-type-not-setting-to-application-json-with-c-json
func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func CreateRouter() *gin.Engine {
	router := gin.Default()

	router.Use(JSONMiddleware())
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
