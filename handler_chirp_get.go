package main

import (
	"net/http"
	"strconv"
	"strings"
)

func (cfg *apiConfig) handlerChirpByID(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/chirps/")
	if path == "" {
		respondWithError(w, http.StatusBadRequest, "Chirp ID is required")
		return
	}

	id, _ := strconv.Atoi(path)

	chirp, err := cfg.DB.GetChirp(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Not found")
		return
	}

	respondWithJSON(w, http.StatusOK, chirp)
}
