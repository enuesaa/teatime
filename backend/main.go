package main

import (
	"io"
	"os"
	"net/http"
	"golang.org/x/net/http2/h2c"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"

	"github.com/joho/godotenv"
	"github.com/enuesaa/teatime-app/backend/gen/v1/v1connect"
)

type SettingsServer struct {
	v1connect.UnimplementedSettingsServiceHandler
}

func main() {
	api := http.NewServeMux()
	api.Handle(v1connect.NewSettingsServiceHandler(&SettingsServer{}))

	mux := http.NewServeMux()
	mux.Handle("/api", http.StripPrefix("/api", api))

	_ := http.ListenAndServe(
		"localhost:80",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
