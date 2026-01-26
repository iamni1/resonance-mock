package main

import (
	httphandlers "Resonance/internal/transport/http"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("GET /api/v1/podcasts", httphandlers.GetPodcastsHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Starting server on :8080")

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
