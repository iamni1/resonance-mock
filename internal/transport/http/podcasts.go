package httphandlers

import (
	"encoding/json"
	"net/http"

	"github.com/iamni1/resonance-mock/internal/service"
)

type Handler struct {
	podcastService service.PodcastService
}

func NewHandler(podcastService service.PodcastService) *Handler {
	return &Handler{
		podcastService: podcastService,
	}
}

func (h *Handler) GetPodcasts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	podcasts, err := h.podcastService.GetAll(ctx)
	if err != nil {
		http.Error(w, "Failed to get podcasts", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(podcasts)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
