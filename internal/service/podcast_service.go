package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamni1/resonance-mock/internal/domain"
)

type PodcastRepository interface {
	GetAll(ctx context.Context) ([]domain.Podcast, error)
	GetByID(ctx context.Context, id string) (domain.Podcast, error)
	Create(ctx context.Context, podcast domain.Podcast) error
}

type PodcastService interface {
	GetAll(ctx context.Context) ([]domain.Podcast, error)
	GetByID(ctx context.Context, id string) (domain.Podcast, error)
	Create(ctx context.Context, input domain.CreatePodcastDTO) (string, error)
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

func (s *podcastService) GetByID(ctx context.Context, id string) (domain.Podcast, error) {

	return s.repo.GetByID(ctx, id)
}

func (s *podcastService) Create(ctx context.Context, input domain.CreatePodcastDTO) (string, error) {
	newID := uuid.NewString()

	podcast := domain.Podcast{
		ID:          newID,
		Title:       input.Title,
		Author:      input.Author,
		CoverURL:    input.CoverURL,
		Description: input.Description,
	}

	err := s.repo.Create(ctx, podcast)
	if err != nil {
		return "", err
	}

	return newID, nil
}
