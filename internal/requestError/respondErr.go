package requesterror

import (
	"net/http"

	responsestype "github.com/ItzTass/Chirpy/internal/responses"
)

func SendreqErr(w http.ResponseWriter, err error, nonRequestString string) {
	if reqErr, ok := err.(*RequestError); ok {
		responsestype.RespondWithErr(w, reqErr.StatusCode, reqErr.Error())
	} else {
		responsestype.RespondWithErr(w, http.StatusInternalServerError, nonRequestString)
	}
}
