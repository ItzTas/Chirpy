package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ItzTass/Chirpy/internal/database"
	"github.com/joho/godotenv"
)

func addDebugFlag() {
	dgb := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	if *dgb {
		fmt.Println("Debug mode enabled")
		err := os.Remove(database_path)
		if err != nil && !os.IsNotExist(err) {
			log.Fatalf("Failed to remove database file: %v", err)
		}
	} else {
		fmt.Println("No debug mode")
	}
}

func main() {
	godotenv.Load()
	jwtSecret := os.Getenv("JWT_SECRET")
	addDebugFlag()
	const filepathRoot = "."
	const port = "8080"
	db, err := database.NewDB(database_path)
	if err != nil {
		log.Fatal(err)
	}
	cfg := apiConfig{
		DB:        db,
		JWTSecret: jwtSecret,
	}

	mux := http.NewServeMux()
	mux.Handle("GET /app/*", http.StripPrefix("/app", cfg.middlewareMetricsInc(http.FileServer(http.Dir(filepathRoot)))))

	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)
	mux.HandleFunc("GET /api/reset", cfg.handlerReset)
	mux.HandleFunc("POST /api/chirps", cfg.handleChirpPost)
	mux.HandleFunc("GET /api/chirps", cfg.handleChirpsGet)
	mux.HandleFunc("GET /api/chirps/{chirpID}", cfg.handleChirpGetByID)
	mux.HandleFunc("POST /api/users", cfg.createUser)
	mux.HandleFunc("POST /api/login", cfg.handleLogin)
	//mux.HandleFunc("PUT /api/users", cfg.handleUsersPUT)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
