package database

import (
	"fmt"
)

func (db *DB) GetUser(email string) (User, error) {
	stru, err := db.loadDB()
	if err != nil {
		return User{}, err
	}
	users := stru.Users
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user with email %s not found", email)
}
