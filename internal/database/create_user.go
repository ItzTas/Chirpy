package database

import (
	"net/http"

	requesterror "github.com/ItzTass/Chirpy/internal/requestError"
)

func (db *DB) CreateUser(email, password string) (User, error) {
	if err := db.ensureDB(); err != nil {
		return User{}, requesterror.NewRequestErr(http.StatusConflict, err.Error())
	}
	if _, exists := db.GetUser(email); exists == nil {
		return User{}, requesterror.NewRequestErr(http.StatusBadRequest, "User already exists")
	}
	dbStruc, err := db.loadDB()
	if err != nil {
		return User{}, requesterror.NewRequestErr(http.StatusInternalServerError, err.Error())
	}
	newID := len(dbStruc.Users) + 1
	newUser := User{
		Id:       newID,
		Email:    email,
		Password: password,
	}
	dbStruc.Users[newID] = newUser
	if err = db.writeDB(dbStruc); err != nil {
		return User{}, requesterror.NewRequestErr(http.StatusInternalServerError, err.Error())
	}
	return newUser, nil
}
