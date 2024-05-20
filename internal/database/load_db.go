package database

import (
	"encoding/json"
	"os"
)

func (db *DB) loadDB() (DBStructure, error) {
	db.mux.Lock()
	defer db.mux.Unlock()
	dat, err := os.ReadFile(db.path)
	if err != nil {
		return DBStructure{}, err
	}

	dbStructure := DBStructure{}
	err = json.Unmarshal(dat, &dbStructure)
	if err != nil {
		return DBStructure{}, err
	}
	return dbStructure, nil
}
