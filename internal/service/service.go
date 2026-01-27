package service

type Service struct {
	PodcastService PodcastService
}

func NewService() *Service {
	podcastService := NewPodcastService()

	return &Service{
		PodcastService: podcastService,
	}
}
