package database

import (
	"encoding/json"
	"os"
)

func (db *DB) writeDB(dbStructure DBStructure) error {
	dat, err := json.Marshal(dbStructure)
	if err != nil {
		return err
	}

	if err = os.WriteFile(db.path, dat, 0644); err != nil {
		return err
	}
	return nil
}
