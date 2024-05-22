package database

import (
	"fmt"
	"net/http"

	requesterror "github.com/ItzTass/Chirpy/internal/requestError"
)

func (db *DB) GetUser(email string) (User, error) {
	stru, err := db.loadDB()
	if err != nil {
		return User{}, requesterror.NewRequestErr(http.StatusInternalServerError, err.Error())
	}
	users := stru.Users
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}
	return User{}, requesterror.NewRequestErr(http.StatusNotFound, fmt.Sprintf("User with email %s not found", email))
}
