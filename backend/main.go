package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// logger
	f, _ := os.Create("tmp/gin.log")
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	router := CreateRouter()
	router.Run(":80")
}
