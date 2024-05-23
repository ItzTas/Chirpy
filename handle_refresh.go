package main

import (
	"net/http"
	"time"

	"github.com/ItzTass/Chirpy/internal/auth"
)

func (cfg *apiConfig) handleRefresh(w http.ResponseWriter, r *http.Request) {
	type returnVals struct {
		Token string `json:"token"`
	}
	reqToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	user, err := cfg.DB.GetUserByRefreshToken(reqToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	now := time.Now().UTC()
	if user.RefreshToken.Expiration.Before(now) || user.RefreshToken.Token == "" || user.RefreshToken.Token == "revoked" {
		respondWithError(w, http.StatusUnauthorized, "Token expired")
		return
	}
	token, err := auth.MakeJWT(user.ID, cfg.jwtSecret, accessTokenDuration)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Could not create token")
		return
	}
	respondWithJSON(w, http.StatusOK, returnVals{
		Token: token,
	})
}
