package main

import (
	"net/http"
	"sort"
	"strconv"
)

const (
	sortTypeAsc  = "asc"
	sortTypeDesc = "desc"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")
	chirpID, err := strconv.Atoi(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID")
		return
	}

	dbChirp, err := cfg.DB.GetChirp(chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp")
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:       dbChirp.ID,
		Body:     dbChirp.Body,
		AuthorID: dbChirp.AuthorID,
	})
}

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	authorID := r.URL.Query().Get("author_id")
	if authorID != "" {
		cfg.chirpsRetrieveFromAuthor(w, r)
		return
	}

	sortType := r.URL.Query().Get("sort")
	if sortType != sortTypeAsc && sortType != sortTypeDesc {
		sortType = sortTypeAsc
	}

	dbChirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps")
		return
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:       dbChirp.ID,
			Body:     dbChirp.Body,
			AuthorID: dbChirp.AuthorID,
		})
	}

	if sortType == sortTypeAsc {
		sort.Slice(chirps, func(i, j int) bool {
			return chirps[i].ID < chirps[j].ID
		})
	} else if sortType == sortTypeDesc {
		sort.Slice(chirps, func(i, j int) bool {
			return chirps[i].ID > chirps[j].ID
		})
	}

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) chirpsRetrieveFromAuthor(w http.ResponseWriter, r *http.Request) {
	authorID, err := strconv.Atoi(r.URL.Query().Get("author_id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid author id")
		return
	}
	authorChirps, err := cfg.DB.GetAuthorChirps(authorID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	sort.Slice(authorChirps, func(i, j int) bool {
		return authorChirps[i].ID < authorChirps[j].ID
	})

	respondWithJSON(w, http.StatusOK, authorChirps)
}
