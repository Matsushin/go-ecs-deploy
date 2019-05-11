package main

import (
	"io"
	"net/http"
)

func mainHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World")
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
}
