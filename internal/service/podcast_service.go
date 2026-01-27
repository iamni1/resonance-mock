package service

import (
	"context"

	"github.com/iamni1/resonance-mock/internal/domain"
)

type PodcastService interface {
	GetAll(ctx context.Context) ([]domain.Podcast, error)
}

type podcastService struct {
}

func NewPodcastService() PodcastService {
	return &podcastService{}
}

func (s *podcastService) GetAll(ctx context.Context) ([]domain.Podcast, error) {
	poscasts := []domain.Podcast{
		{
			ID:       "1d7c3563-72c2-4a11-8c3b-1b7f4e9a8f2a",
			Title:    "Go-шные посиделки",
			Author:   "Техлид и Стажер",
			CoverURL: "https://example.com/cover1.jpg",
		},
		{
			ID:       "f47ac10b-58cc-4372-a567-0e02b2c3d479",
			Title:    "Kotlin для самых маленьких",
			Author:   "Подруга",
			CoverURL: "https://example.com/cover2.jpg",
		},
	}

	return poscasts, nil
}
