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
	response := createResponse(r)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(response))
}

func createResponse(r *http.Request) string {
	method := r.Method
	url := r.URL
	query := r.URL.Query()

	return fmt.Sprintf("%s %s %s", method, url, query)
}
