package main

import (
	"fmt"
	"net/http"
)

func mainHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, World！！\n")
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
}
