package bookshelf

import (
	"net/http"

	"github.com/enuesaa/teatime-app/backend/services/books"
	"github.com/gin-gonic/gin"
)

func ListBooks(ctx *gin.Context) {
	list := books.ListBooks()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello alworld",
		"books": list,
	})
}