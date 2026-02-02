package repository

import (
	"context"
	"fmt"

	"github.com/iamni1/resonance-mock/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type podcastRepo struct {
	pool *pgxpool.Pool
}

func NewPodcastRepository(pool *pgxpool.Pool) *podcastRepo {
	return &podcastRepo{
		pool: pool,
	}
}

func (r *podcastRepo) GetAll(ctx context.Context) ([]domain.Podcast, error) {

	query := `SELECT id, title, author, cover_url FROM podcasts`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	defer rows.Close()

	var podcasts []domain.Podcast

	for rows.Next() {
		var p domain.Podcast

		err := rows.Scan(&p.ID, &p.Title, &p.Author, &p.CoverURL)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		podcasts = append(podcasts, p)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return podcasts, nil
}

func (r *podcastRepo) GetByID(ctx context.Context, id string) (domain.Podcast, error) {

	query := `SELECT id, title, author, cover_url, description FROM podcasts WHERE id = $1`

	var p domain.Podcast

	err := r.pool.QueryRow(ctx, query, id).Scan(&p.ID, &p.Title, &p.Author, &p.CoverURL, &p.Description)
	if err != nil {
		return domain.Podcast{}, fmt.Errorf("failed to scan row: %w", err)
	}

	return p, nil
}

func (r *podcastRepo) Create(ctx context.Context, podcast domain.Podcast) error {
	query := `INSERT INTO podcasts (id, title, author, cover_url, description) 
			   VALUES ($1, $2, $3, $4, $5)`

	_, err := r.pool.Exec(ctx, query,
		podcast.ID,
		podcast.Title,
		podcast.Author,
		podcast.CoverURL,
		podcast.Description,
	)

	if err != nil {
		return fmt.Errorf("failed to create podcast: %w", err)
	}

	return nil
}
