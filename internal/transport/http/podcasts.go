package http

import (
	"Resonance/internal/domain"
	"encoding/json"
	"net/http"
)

func GetPodcastsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	podcasts := []domain.Podcast{
		{
			ID:       "1d7c3563-72c2-4a11-8c3b-1b7f4e9a8f2a",
			Title:    "Go-шные посиделки",
			Author:   "Техлид и Стажер",
			CoverURL: "https://example.com/cover1.jpg",
		},
		{
			ID:       "f47ac10b-58cc-4372-a567-0e02b2c3d479",
			Title:    "Kotlin для самых маленьких",
			Author:   "Андроид орео",
			CoverURL: "https://example.com/cover2.jpg",
		},
	}

	if err := json.NewEncoder(w).Encode(podcasts); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
