package service

import (
	"context"

	"github.com/iamni1/resonance-mock/internal/domain"
)

type PodcastRepository interface {
	GetAll(ctx context.Context) ([]domain.Podcast, error)
}

type PodcastService interface {
	GetAll(ctx context.Context) ([]domain.Podcast, error)
}

type podcastService struct {
	repo PodcastRepository
}

func NewPodcastService(repo PodcastRepository) PodcastService {
	return &podcastService{
		repo: repo,
	}
}

func (s *podcastService) GetAll(ctx context.Context) ([]domain.Podcast, error) {

	return s.repo.GetAll(ctx)
}
