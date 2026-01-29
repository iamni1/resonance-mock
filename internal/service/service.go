package service

import (
	"github.com/iamni1/resonance-mock/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	PodcastService PodcastService
}

func NewService(pool *pgxpool.Pool) *Service {
	repo := repository.NewPodcastRepository(pool)
	podcastService := NewPodcastService(repo)

	return &Service{
		PodcastService: podcastService,
	}
}
