package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/enuesaa/teatime-app/backend/binding"
)

type BoardController struct {}
func (ctl *BoardController) Add (c *gin.Context) {
	var body v1.AddBoardRequest
	if !binding.Validate(c, &body) { return }
	c.JSON(200, v1.AddBoardResponse {})
}
