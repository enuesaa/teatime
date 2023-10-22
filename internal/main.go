// teatime-app backend
package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	f, _ := os.Create("tmp/gin.log")
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	router := setupRouter()
	router.Run(":80")
}
