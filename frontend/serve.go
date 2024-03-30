package frontend

import (
	"embed"
	"fmt"
	"mime"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

//go:generate pnpm install
//go:generate pnpm build

//go:embed all:dist/*
var dist embed.FS

func Serve(c echo.Context) error {
	path := c.Request().URL.Path // like `/`
	path = fmt.Sprintf("dist%s", path)
	if strings.HasSuffix(path, "/") {
		path += "index.html"
	}

	f, err := dist.ReadFile(path)
	if err != nil {
		return err
	}
	fileExt := filepath.Ext(path)
	mimeType := mime.TypeByExtension(fileExt)

	return c.Blob(200, mimeType, f)
}
