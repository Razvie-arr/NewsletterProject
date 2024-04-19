package repository

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"newsletterProject/pkg/id"
	dbmodel "newsletterProject/repository/sql/model"
	"newsletterProject/repository/sql/mutation"
	"newsletterProject/repository/sql/query"
	"newsletterProject/service/model"
	"newsletterProject/transport/util"
)

type SubscriberRepository struct {
	pool *pgxpool.Pool
}

func NewSubscriberRepository(pool *pgxpool.Pool) *SubscriberRepository {
	return &SubscriberRepository{
		pool: pool,
	}
}

func (r *SubscriberRepository) ReadSubscriberByEmail(ctx context.Context, email string) (*model.Subscriber, error) {
	var subscriber dbmodel.Subscriber

	if err := pgxscan.Get(
		ctx,
		r.pool,
		&subscriber,
		query.ReadSubscriberByEmail,
		pgx.NamedArgs{
			"email": email,
		},
	); err != nil {
		return nil, err
	}

	return &model.Subscriber{
		ID:    subscriber.Id,
		Email: subscriber.Email,
	}, nil
}

func (r *SubscriberRepository) CreateSubscriber(ctx context.Context, email string) (*model.Subscriber, error) {
	var subscriber dbmodel.Subscriber

	err := r.pool.QueryRow(
		ctx,
		mutation.CreateSubscriber,
		pgx.NamedArgs{
			"email": email,
		},
	).Scan(&subscriber.Id, &subscriber.Email)

	if err != nil {
		return nil, err
	}

	return &model.Subscriber{
		ID:    subscriber.Id,
		Email: subscriber.Email,
	}, nil
}

func (r *SubscriberRepository) Subscribe(ctx context.Context, newsletterId id.ID, subscriberId id.ID) (string, error) {
	// Generate verification string (random string length 32)
	verificationString := util.GenerateRandomString(32)

	// Insert the subscription
	_, err := r.pool.Exec(
		ctx,
		mutation.Subscribe,
		pgx.NamedArgs{
			"newsletter_id":       newsletterId.String(),
			"subscriber_id":       subscriberId.String(),
			"verification_string": verificationString,
		},
	)
	if err != nil {
		return "", err
	}

	return verificationString, nil
}

func (r *SubscriberRepository) GetVerificationString(ctx context.Context, newsletterId id.ID, subscriberId id.ID) (string, error) {
	var verificationString string

	if err := pgxscan.Get(
		ctx,
		r.pool,
		&verificationString,
		query.GetVerificationString,
		pgx.NamedArgs{
			"newsletter_id": newsletterId.String(),
			"subscriber_id": subscriberId.String(),
		},
	); err != nil {
		return "", err
	}

	return verificationString, nil
}

func (r *SubscriberRepository) Unsubscribe(ctx context.Context, newsletterId id.ID, subscriberId id.ID) error {
	_, err := r.pool.Exec(
		ctx,
		mutation.Unsubscribe,
		pgx.NamedArgs{
			"newsletter_id": newsletterId.String(),
			"subscriber_id": subscriberId.String(),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
