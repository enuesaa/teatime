package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/enuesaa/teatime-app/backend/binding"
)

func AddBoard(c *gin.Context) {
	var body v1.AddBoardRequest
	if !binding.Validate(c, &body) {
		return
	}
	
	c.JSON(200, v1.AddBoardResponse {})
}
