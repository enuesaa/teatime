package bookmark

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/enuesaa/teatime-app/backend/validate"
)

func AddBookmark(c *gin.Context) {
	var body v1.AddBookmarkRequest
	if err := validate.Validate(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, v1.AddBookmarkResponse {})
}
