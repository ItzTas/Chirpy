package database

func (db *DB) GetChirps() ([]Chirp, error) {
	db.mux.Lock()
	defer db.mux.Unlock()
	dbStructure, err := db.loadDB()
	if err != nil {
		return []Chirp{}, err
	}
	chirps := []Chirp{}
	for _, chirp := range dbStructure.Chirps {
		chirps = append(chirps, chirp)
	}
	return chirps, nil
}
