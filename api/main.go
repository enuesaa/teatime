package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    method := r.Method
    url := r.URL
    query := r.URL.Query()

    fmt.Fprintf(w, "%s %s\n", method, url)
    if query == nil {
        return
    }
    for key, val := range query {
        fmt.Fprintf(w, "%s = %s\n", key, val[0])
    }
}
