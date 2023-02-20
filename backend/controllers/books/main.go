package books

import (
	"net/http"

	"github.com/enuesaa/teatime-app/backend/services/books"
	"github.com/gin-gonic/gin"
)

type Books struct {
    Message   string    `json:"message" example:"aa"`
}
// ListBooks list books
// @Summary      List Books
// @Description  
// @Accept       json
// @Produce      json
// @Success      200  {object}  Books
// @Router       /books/list [get]
func ListBooks(ctx *gin.Context) {
	list := books.ListBooks()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello alworld",
		"books": list,
	})
}