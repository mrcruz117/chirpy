package main

import (
	"encoding/json"
	"net/http"

	"github.com/mrcruz117/chirpy/internal/database"
)

func (cfg *apiConfig) handlerUserLogin(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := cfg.DB.UserAuth(database.User{
		Email:    params.Email,
		Password: params.Password,
	})
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Not Authorized")
		return
	}

	respondWithJSON(w, http.StatusOK, User{
		ID:    user.ID,
		Email: user.Email,
	})
}
