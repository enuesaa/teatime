package bookmark

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
)

func validateAddBookmark(body *v1.AddBookmarkRequest) bool {
	return body.Name != ""
}

func AddBookmark(ctx *gin.Context) {
	// see https://gin-gonic.com/ja/docs/examples/binding-and-validation/
	var body v1.AddBookmarkRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if !validateAddBookmark(&body) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello alaworld",
		"redisvalue": body.Name,
	})
}
