package main

import (
	"log"
	"net/http"
	"sync/atomic"

	// "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	jwtSecret      string
	polkaKey       string
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

	// keep pointer or not?
	apiCfg := &apiConfig{}
	fileServer := http.FileServer(http.Dir(filepathRoot))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", fileServer)))
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("POST /api/validate_chirp", handlerChirpsValidate)

	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())

}
