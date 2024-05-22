package main

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (cfg *apiConfig) createJWT(id, expiresInSeconds int) (string, error) {
	if expiresInSeconds == 0 || expiresInSeconds > 86400 {
		expiresInSeconds = 86400
	}

	claims := jwt.RegisteredClaims{
		Issuer:    "Chirpy",
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Duration(expiresInSeconds) * time.Second)),
		Subject:   strconv.Itoa(id),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
