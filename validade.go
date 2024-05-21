package main

import (
	"net/http"
)

// func handlerValidade(w http.ResponseWriter, r *http.Request) {
// 	type paramethers struct {
// 		Body string `json:"body"`
// 	}
// 	type returnVals struct {
// 		Cleaned_Blody string `json:"cleaned_body"`
// 	}
// 	decoder := json.NewDecoder(r.Body)
// 	params := paramethers{}
// 	err := decoder.Decode(&params)
// 	if err != nil {
// 		respondWithErr(w, http.StatusInternalServerError, "Something went wrong")
// 		return
// 	}

// 	if !validadeMaxLenght(w, len(params.Body)) {
// 		return
// 	}

// 	cleaned := validadeProfane(params.Body)

// 	respondWithJSON(w, http.StatusOK, returnVals{
// 		Cleaned_Blody: cleaned,
// 	})
// }

func validadeMaxLenght(w http.ResponseWriter, leng int) bool {
	const maxChirpLength = 140
	if leng > maxChirpLength {
		respondWithErr(w, http.StatusBadRequest, "Chirp is too long")
		return false
	}
	return true
}

// func validadeProfane(toV string) string {
// 	words := strings.Split(toV, " ")
// 	profaneWords := []string{"kerfuffle", "sharbert", "fornax"}
// 	for i, word := range words {
// 		if slices.Contains(profaneWords, strings.ToLower(word)) {
// 			words[i] = "****"
// 		}
// 	}
// 	return strings.Join(words, " ")
// }
