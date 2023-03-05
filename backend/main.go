package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/enuesaa/teatime-app/backend/controllers/setting"
	"github.com/enuesaa/teatime-app/backend/controllers/lounge"
	"github.com/enuesaa/teatime-app/backend/controllers/bookshelf"
)

func jsonMiddleware() gin.HandlerFunc {
	// https://stackoverflow.com/questions/41109065/golang-gin-gonic-content-type-not-setting-to-application-json-with-c-json
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func main() {
	// logger
	f, _ := os.Create("tmp/gin.log")
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	router := gin.Default()
	router.Use(jsonMiddleware())
	base := router.Group("/api")
	{
		settingRoute := base.Group("/v1.SettingService")
		settingRoute.POST("/GetAppearance", setting.GetAppearance)
		
		loungeRoute := base.Group("/v1.LoungeService")
		loungeRoute.POST("/Spotify", lounge.CallSpotifyApi)

		bookshelfRoute := base.Group("/v1.BookShelfService")
		bookshelfRoute.POST("/List", bookshelf.ListBooks)
	}
	router.Run(":80")
}
