package repository

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"newsletterProject/pkg/id"
	dbmodel "newsletterProject/repository/sql/model"
	"newsletterProject/repository/sql/query"
	"newsletterProject/service/model"
)

type NewsletterRepository struct {
	pool *pgxpool.Pool
}

func NewNewsletterRepository(pool *pgxpool.Pool) *NewsletterRepository {
	return &NewsletterRepository{
		pool: pool,
	}
}

func (r *NewsletterRepository) ReadNewsletter(ctx context.Context, newsletterId id.ID) (*model.Newsletter, error) {
	var newsletter dbmodel.Newsletter

	if err := pgxscan.Get(
		ctx,
		r.pool,
		&newsletter,
		query.ReadNewsletter,
		pgx.NamedArgs{
			"id": newsletterId,
		},
	); err != nil {
		return nil, err
	}

	var editor dbmodel.Editor

	// Get editor
	if err := pgxscan.Get(
		ctx,
		r.pool,
		&editor,
		query.ReadEditor,
		pgx.NamedArgs{
			"id": newsletter.EditorId,
		},
	); err != nil {
		return nil, err
	}

	// Get subscribers
	var subscribers []dbmodel.Subscriber

	if err := pgxscan.Select(
		ctx,
		r.pool,
		&subscribers,
		query.ReadSubscribersByNewsletterId,
		pgx.NamedArgs{
			"newsletter_id": newsletterId,
		},
	); err != nil {
		return nil, err
	}

	// Convert subscribers to model
	var svcSubscribers []model.Subscriber
	for _, subscriber := range subscribers {
		svcSubscribers = append(svcSubscribers, model.Subscriber{
			ID:    subscriber.Id,
			Email: subscriber.Email,
		})
	}

	var description *string
	if newsletter.Description.Valid {
		description = &newsletter.Description.String
	}

	return &model.Newsletter{
		ID:          newsletter.Id,
		Name:        newsletter.Name,
		Description: description,
		Editor: model.Editor{
			ID:    editor.Id,
			Email: editor.Email,
		},
		Subscribers: svcSubscribers,
	}, nil

}
