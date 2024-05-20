package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"
	cfg := apiConfig{}

	mux := http.NewServeMux()
	mux.Handle("GET /app/*", http.StripPrefix("/app", cfg.middlewareMetricsInc(http.FileServer(http.Dir(filepathRoot)))))

	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)
	mux.HandleFunc("GET /api/reset", cfg.handlerReset)
	// mux.HandleFunc("POST /api/chirps")

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
