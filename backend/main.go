package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

// @title テスト
// @version 0.1.0
// @description テスト
// @host http://localhost:3000/api
func main() {
	// logger
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	router := gin.Default()
	router.GET("/api", func(ctx *gin.Context) {
		spotifyClientId := os.Getenv("SPOTIFY_CLIENT_ID")
		spotifyClientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

		config := &clientcredentials.Config{
			ClientID:     spotifyClientId,
			ClientSecret: spotifyClientSecret,
			TokenURL:     spotifyauth.TokenURL,
		}
		token, err := config.Token(ctx)
		if err != nil {
			log.Fatalf("couldn't get token: %v", err)
		}

		// https://zenn.dev/shimpo/articles/trying-spotify-api-with-go
		httpClient := spotifyauth.New().Client(ctx, token)
		client := spotify.New(httpClient)

		result, err := client.Search(ctx, "Beatles", spotify.SearchTypeArtist, spotify.Limit(2))
		if err != nil {
			log.Fatal(err)
		}

		bytes, _ := json.MarshalIndent(result, "", "  ")
		fmt.Fprintln(gin.DefaultWriter, string(bytes))

		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello alworld",
			"data":    string(bytes),
		})
	})
	router.Run(":80")
}
