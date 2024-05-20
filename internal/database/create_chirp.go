package database

func (db *DB) CreateChirp(body string) (Chirp, error) {
	if err := db.ensureDB(); err != nil {
		return Chirp{}, err
	}
	dbstru, err := db.loadDB()
	if err != nil {
		return Chirp{}, err
	}
	newID := 1
	for id := range dbstru.Chirps {
		if id >= newID {
			newID = id + 1
		}
	}
	newC := Chirp{
		Id:   newID,
		Body: body,
	}
	dbstru.Chirps[newID] = newC
	if err = db.writeDB(dbstru); err != nil {
		return Chirp{}, err
	}
	return newC, nil
}
