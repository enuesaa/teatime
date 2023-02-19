package books

import (
    "io"
	"net/http"
)

func ListBooks () string {
	// see https://tutuz-tech.hatenablog.com/entry/2020/03/22/160529
	res, err := http.Get("https://www.googleapis.com/books/v1/volumes?q=ハンズオン")
    if err != nil {
        panic(err)
    }
	defer res.Body.Close()
    body, _ := io.ReadAll(res.Body)
	return string(body)
}