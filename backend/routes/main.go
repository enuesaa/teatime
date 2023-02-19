package routes

import (
	"github.com/enuesaa/teatime-app/backend/controllers/settings"
	"github.com/enuesaa/teatime-app/backend/controllers/musics"
	"github.com/enuesaa/teatime-app/backend/controllers/books"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	base := router.Group("/api")
	{
		settingsRoute := base.Group("/settings")
		settingsRoute.GET("/appearance", settings.GetAppearance)
		settingsRoute.PUT("/appearance", settings.PutAppearance)
		
		musicsRoute := base.Group("/musics")
		musicsRoute.GET("/spotify", musics.Spotify)

		booksRoute := base.Group("/books")
		booksRoute.GET("/list", books.ListBooks)
	}

	return router
}
