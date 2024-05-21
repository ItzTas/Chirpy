package database

func (db *DB) CreateUser(email, password string) (User, error) {
	if err := db.ensureDB(); err != nil {
		return User{}, err
	}
	dbStruc, err := db.loadDB()
	if err != nil {
		return User{}, err
	}
	newID := 1
	for id := range dbStruc.Users {
		if id >= newID {
			newID = id + 1
		}
	}
	newUser := User{
		Id:       newID,
		Email:    email,
		Password: password,
	}
	dbStruc.Users[newID] = newUser
	if err = db.writeDB(dbStruc); err != nil {
		return User{}, err
	}
	return newUser, nil
}
