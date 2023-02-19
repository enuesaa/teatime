package musics

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

// see https://stackoverflow.com/questions/70208440/how-to-disable-golang-unused-function-error
//lint:ignore U1000 Ignore unused function
func callSpotifyApi(ctx *gin.Context) {
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
}
