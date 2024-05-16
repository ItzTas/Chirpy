package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Handler: mux,
		Addr:    "localhost:8080",
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
