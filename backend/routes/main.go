package routes

import (
	"github.com/enuesaa/teatime-app/backend/controllers/musics"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()
	base := router.Group("/api")
	{
		musicsRoute := base.Group("/musics")
		musicsRoute.GET("/spotify", musics.Spotify)
	}

	return router
}
