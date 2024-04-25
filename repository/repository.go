package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	*EditorRepository
	*SubscriberRepository
	*NewsletterRepository
	*PostRepository
}

func New(pool *pgxpool.Pool) (*Repository, error) {
	newsletterRepository := NewNewsletterRepository(pool)
	return &Repository{
		EditorRepository:     NewEditorRepository(pool),
		SubscriberRepository: NewSubscriberRepository(pool),
		NewsletterRepository: newsletterRepository,
		PostRepository:       NewPostRepository(pool, newsletterRepository),
	}, nil
}
