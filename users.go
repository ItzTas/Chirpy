package main

import (
	"encoding/json"
	"net/http"

	requesterror "github.com/ItzTass/Chirpy/internal/requestError"
	responsestype "github.com/ItzTass/Chirpy/internal/responses"
	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type returnVals struct {
		Email string `json:"email"`
		Id    int    `json:"id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		requesterror.SendreqErr(w, err, "Could not decode params")
		return
	}

	if !validadeMaxLenght(w, len(params.Email)) {
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		requesterror.SendreqErr(w, err, "Internal Server Error")
	}

	user, err := cfg.DB.CreateUser(params.Email, string(hashed))
	if err != nil {
		requesterror.SendreqErr(w, err, "Internal Server Error")
		return
	}

	responsestype.RespondWithJSON(w, http.StatusCreated, returnVals{
		Email: user.Email,
		Id:    user.Id,
	})
}

func (cfg *apiConfig) handleLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password           string `json:"password"`
		Email              string `json:"email"`
		Expires_in_seconds int    `json:"expires_in_seconds"`
	}
	type returnVals struct {
		Email string `json:"email"`
		Id    int    `json:"id"`
		Token string `json:"token"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		requesterror.SendreqErr(w, err, "Could not decode params")
		return
	}
	user, err := cfg.DB.GetUser(params.Email)
	if err != nil {
		requesterror.SendreqErr(w, err, "Internal Server Error")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		requesterror.SendreqErr(w, err, "Internal Server Error")
		return
	}
	token, err := cfg.createJWT(user.Id, params.Expires_in_seconds)
	if err != nil {
		requesterror.SendreqErr(w, err, "Internal Server Error")
		return
	}
	responsestype.RespondWithJSON(w, http.StatusOK, returnVals{
		Email: user.Email,
		Id:    user.Id,
		Token: token,
	})
}

// func (cfg *apiConfig) handleUsersPUT(w http.ResponseWriter, r *http.Request) {
// 	token, err := getToken(r)
// 	if err != nil {
// 		respondWithErr(w, http.StatusInternalServerError, "Could not create a token")
// 		return
// 	}
// }

// func getToken(r *http.Request) (string, error) {
// 	auto := r.Header.Get("Authorization")
// 	if auto == "" {
// 		return "", requesterror.NewRequestErr(http.StatusBadRequest, "invalid autorization")
// 	}

// 	parts := strings.Split(auto, " ")
// 	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
// 		return "", requesterror.NewRequestErr(http.StatusBadRequest, "invalid authorization header format")
// 	}

// 	return parts[1], nil
// }
