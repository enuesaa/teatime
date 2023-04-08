package main

import (
	"github.com/enuesaa/teatime-app/backend/controller/board"
	"github.com/enuesaa/teatime-app/backend/controller/bookmark"
	"github.com/enuesaa/teatime-app/backend/controller/feed"
	"github.com/enuesaa/teatime-app/backend/controller/setting"
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
		settingCtl := setting.SettingController{}
		settingRoute.POST("/GetAppearance", settingCtl.Get)
		settingRoute.POST("/PutAppearance", settingCtl.Put)

		bookmarkRoute := base.Group("/v1.Bookmark")
		bookmarkCtl := bookmark.BookmarkController{}
		bookmarkRoute.POST("/ListBookmarks", bookmarkCtl.List)
		bookmarkRoute.POST("/GetBookmark", bookmarkCtl.Get)
		bookmarkRoute.POST("/AddBookmark", bookmarkCtl.Add)
		bookmarkRoute.POST("/UpdateBookmark", bookmarkCtl.Update)
		bookmarkRoute.POST("/DeleteBookmark", bookmarkCtl.Delete)

		feedRoute := base.Group("/v1.Feed")
		feedCtl := feed.FeedController{}
		feedRoute.POST("/ListFeeds", feedCtl.List)
		feedRoute.POST("/AddFeed", feedCtl.Add)
		feedRoute.POST("/GetFeed", feedCtl.Get)
		feedRoute.POST("/ListItems", feedCtl.ListItems)
		feedRoute.POST("/GetAppearance", feedCtl.GetAppearance)
		feedRoute.POST("/UpdateAppearance", feedCtl.UpdateAppearance)
		feedRoute.POST("/Fetch", feedCtl.Fetch)
		feedRoute.POST("/DeleteFeed", feedCtl.Delete)

		boardRoute := base.Group("/v1.Board")
		boardCtl := board.BoardController{}
		boardRoute.POST("/ListBoards", boardCtl.List)
		boardRoute.POST("/GetBoard", boardCtl.Get)
		boardRoute.POST("/AddBoard", boardCtl.Add)
		boardRoute.POST("/UpdateBoard", boardCtl.Update)
		boardRoute.POST("/DeleteBoard", boardCtl.Delete)
		boardRoute.POST("/Checkin", boardCtl.Checkin)
		boardRoute.POST("/ListTimeline", boardCtl.ListTimeline)
		boardRoute.POST("/Archive", boardCtl.Archive)
		boardRoute.POST("/UnArchive", boardCtl.UnArchive)
	}

	return router
}
