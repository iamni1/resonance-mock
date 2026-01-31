package httphandlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/iamni1/resonance-mock/internal/service"
	"github.com/jackc/pgx/v5"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) GetPodcasts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	podcasts, err := h.services.PodcastService.GetAll(ctx)
	if err != nil {
		http.Error(w, "Failed to get podcasts", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(podcasts)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *Handler) GetPodcastByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	_, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid podcast ID format", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	podcast, err := h.services.PodcastService.GetByID(ctx, idStr)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Podcast not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(podcast); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
