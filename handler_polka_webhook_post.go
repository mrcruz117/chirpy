package main

import (
	"encoding/json"
	"net/http"

	"github.com/mrcruz117/chirpy/internal/auth"
	// "github.com/mrcruz117/chirpy/internal/database"
)

func (cfg *apiConfig) handlerWebhook(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID int `json:"user_id"`
		} `json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	apiKeyErr := auth.GetApiKey(r.Header)
	if apiKeyErr != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find API key")
		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	user, err := cfg.DB.GetUser(params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	user.IsChirpyRed = true
	_, err = cfg.DB.UpgradeChirpyRed(params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't update user")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
