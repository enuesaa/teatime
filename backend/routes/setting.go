package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SettingController struct{}

func (ctrl SettingController) One(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "ab"})
}

func (ctrl SettingController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "a"})
}
