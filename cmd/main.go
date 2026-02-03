package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/iamni1/resonance-mock/internal/config"
	"github.com/iamni1/resonance-mock/internal/pkg/postgres"
	"github.com/iamni1/resonance-mock/internal/service"
	httphandlers "github.com/iamni1/resonance-mock/internal/transport/http"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	slog.Info("Starting application", "port", 8080, "env", "local")

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Config error: ", err)
	}

	ctx := context.Background()

	pool, err := postgres.NewClient(ctx, cfg.DB_DSN)
	if err != nil {
		log.Fatal("DB init error: ", err)
	}
	defer pool.Close()

	services := service.NewService(pool)
	handler := httphandlers.NewHandler(services)

	mux := http.NewServeMux()


	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("OK"))
	})
	mux.HandleFunc("GET /api/v1/podcasts", handler.GetPodcasts)
	mux.HandleFunc("GET /api/v1/podcasts/{id}", handler.GetPodcastByID)
	mux.HandleFunc("POST /api/v1/podcasts", handler.CreatePodcast)

	wrappedMux := httphandlers.LoggingMiddleware(mux)

	
	server := http.Server{
		Addr:    ":8080",
		Handler: wrappedMux,
	}

	log.Println("Starting server on :8080")

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
