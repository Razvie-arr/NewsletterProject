package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
}

func New(pool *pgxpool.Pool) (*Repository, error) {
	return &Repository{}, nil
}
