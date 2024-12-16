package main

import (
	"log"
	"net/http"

	// "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	fileserverHits int
	// DB             *database.DB
	jwtSecret string
	polkaKey  string
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// jwtSecret := os.Getenv("JWT_SECRET")
	// polkaKey := os.Getenv("POLKA_KEY")
	// if jwtSecret == "" {
	// 	log.Fatal("JWT_SECRET environment variable is not set")
	// }
	const filepathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())

	// always returns a 404

}
