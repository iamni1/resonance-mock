package main

import (
	"log"
	"net/http"

	"github.com/iamni1/resonance-mock/internal/service"
	httphandlers "github.com/iamni1/resonance-mock/internal/transport/http"
)

func main() {

	mux := http.NewServeMux()

	podcastService := service.NewPodcastService()

	mainHandler := httphandlers.NewHandler(podcastService)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("GET /api/v1/podcasts", mainHandler.GetPodcasts)

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
