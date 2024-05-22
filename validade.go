package main

import (
	"net/http"

	responsestype "github.com/ItzTass/Chirpy/internal/responses"
)

func validadeMaxLenght(w http.ResponseWriter, leng int) bool {
	const maxChirpLength = 140
	if leng > maxChirpLength {
		responsestype.RespondWithErr(w, http.StatusBadRequest, "Chirp is too long")
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
