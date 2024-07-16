package main

// server init
/*
Create a new http.ServeMux
Create a new http.Server struct and use the new "ServeMux" as the server's handler
Use the server's ListenAndServe method to start the server
*/

import (
	"fmt"
	"net/http"
)

func main() {
	// create a new ServeMux
	mux := http.NewServeMux()

	// create a new server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// start the server
	fmt.Println("Starting server on port 8080")
	server.ListenAndServe()
}
