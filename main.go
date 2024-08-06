package main

import (
	"fmt"
	"net/http"
)

func main() {
	// create a new ServeMux
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("."))

	mux.Handle("/", fileServer)
	// create a new server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// start the server
	fmt.Println("Starting server on port 8080")
	server.ListenAndServe()
}
