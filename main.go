package main

import (
	"fmt"
	"net/http"
)

func mainHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Deploy succeeded.\n")
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
}
