package domain

type CreatePodcastDTO struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	CoverURL    string `json:"cover_url"`
	Description string `json:"description"`
}
