package main

import (
	"io"
	"os"

	"github.com/enuesaa/teatime-app/backend/routes"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

// @title テスト
// @version 0.1.0
// @description テスト
// @host http://localhost:3000/api
func main() {
	// logger
	f, _ := os.Create("tmp/gin.log")
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	// @todo handle in dockerfile
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	router := routes.CreateRouter()
	router.Run(":80")
}
