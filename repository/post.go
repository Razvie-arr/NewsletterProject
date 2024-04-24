package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"newsletterProject/pkg/id"
	dbmodel "newsletterProject/repository/sql/model"
	"newsletterProject/repository/sql/mutation"
	"newsletterProject/service/model"
)

type PostRepository struct {
	pool                 *pgxpool.Pool
	newsletterRepository *NewsletterRepository
}

func NewPostRepository(pool *pgxpool.Pool, newsletterRepository *NewsletterRepository) *PostRepository {
	return &PostRepository{pool, newsletterRepository}
}

func (r *PostRepository) CreatePost(ctx context.Context, content string, newsletterId id.ID) (*model.Post, error) {
	var post dbmodel.Post

	// Create and get the post
	if err := r.pool.QueryRow(
		ctx,
		mutation.CreatePost,
		pgx.NamedArgs{
			"content":       content,
			"newsletter_id": newsletterId.String(),
		}).Scan(&post.Id, &post.Content, &post.NewsletterId); err != nil {
		return nil, err
	}

	newsletter, err := r.newsletterRepository.ReadNewsletter(ctx, newsletterId)
	if err != nil {
		return nil, err
	}

	return &model.Post{
		ID:         post.Id,
		Content:    post.Content,
		Newsletter: *newsletter,
	}, nil
}
