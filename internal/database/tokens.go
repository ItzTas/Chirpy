package database

import (
	"time"

	"github.com/ItzTass/Chirpy/internal/auth"
)

type RefreshToken struct {
	Token      string
	Expiration time.Time
}

func (db *DB) CreateAndStoreRefTokenToBD(userID int, expiration time.Duration) (RefreshToken, error) {
	token, err := auth.GenerateRefreshToken()
	if err != nil {
		return RefreshToken{}, err
	}
	refToken := RefreshToken{
		Token:      token,
		Expiration: time.Now().UTC().Add(expiration),
	}

	dbStructure, err := db.loadDB()
	if err != nil {
		return RefreshToken{}, err
	}
	user, ok := dbStructure.Users[userID]
	if !ok {
		return RefreshToken{}, ErrNotExist
	}
	user.RefreshToken = refToken
	dbStructure.Users[userID] = user
	if err := db.writeDB(dbStructure); err != nil {
		return RefreshToken{}, err
	}
	return refToken, nil
}

func (db *DB) RevokeToken(refreshTokenString string) error {
	user, err := db.GetUserByRefreshToken(refreshTokenString)
	if err != nil {
		return err
	}
	dbStructure, err := db.loadDB()
	if err != nil {
		return err
	}
	user, ok := dbStructure.Users[user.ID]
	if !ok {
		return ErrNotExist
	}
	user.RefreshToken.Token = "revoked"
	dbStructure.Users[user.ID] = user
	err = db.writeDB(dbStructure)
	if err != nil {
		return err
	}
	return nil
}
