package main

import (
	"fmt"
	"net/http"

	"github.com/ItzTass/Chirpy/internal/database"
)

type apiConfig struct {
	emails      []string
	fileservers int
	DB          *database.DB
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileservers++
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "text/html")
	w.Write([]byte(fmt.Sprintf(`
	<html>
	
	<body>
		<h1>Welcome, Chirpy Admin</h1>
		<p>Chirpy has been visited %d times!</p>
	</body>
	
	</html>
	`, cfg.fileservers)))
}

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileservers = 0
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reseted to 0"))
}
