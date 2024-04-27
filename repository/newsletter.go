package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"newsletterProject/pkg/id"
	dbmodel "newsletterProject/repository/sql/model"
	"newsletterProject/repository/sql/mutation"
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

func (r *NewsletterRepository) CreateNewsletter(ctx context.Context, name, description string, editorId id.ID) (*model.BaseNewsletter, error) {
	var newsletter dbmodel.Newsletter

	namedArgs := pgx.NamedArgs{
		"name":      name,
		"editor_id": editorId,
	}

	if description != "" {
		namedArgs["description"] = description
	} else {
		namedArgs["description"] = nil
	}

	if err := r.pool.QueryRow(
		ctx,
		mutation.CreateNewsletter,
		namedArgs,
	).Scan(&newsletter.Id, &newsletter.Name, &newsletter.Description); err != nil {
		return nil, errors.New("Error inserting editor to DB: " + err.Error())
	}

	var nullableDescription *string
	if newsletter.Description.Valid {
		nullableDescription = &newsletter.Description.String
	}

	return &model.BaseNewsletter{
		ID:          newsletter.Id,
		Name:        newsletter.Name,
		Description: nullableDescription,
	}, nil
}

func (r *NewsletterRepository) ExistsNewsletterWithEditor(ctx context.Context, newsletterId id.ID, editorId id.ID) error {
	var exists bool

	if err := r.pool.QueryRow(
		ctx,
		query.ExistsNewsletterWithEditor,
		pgx.NamedArgs{
			"id":       newsletterId.String(),
			"editorId": editorId.String(),
		},
	).Scan(&exists); err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("no newsletter found with the specified editor and newsletter ID")
	}

	return nil
}

func (r *NewsletterRepository) DeleteNewsletter(ctx context.Context, newsletterId id.ID) error {
	// pouze mažeme, protože již máme ověřené, že newsletter existuje a patří danému editorovi
	_, err := r.pool.Exec(
		ctx,
		mutation.DeleteNewsletter,
		pgx.NamedArgs{
			"id": newsletterId.String(),
		},
	)
	return err
}

func (r *NewsletterRepository) GetNewslettersInfo(ctx context.Context, lim int) ([]*model.NewsletterInfo, error) {
	var dbNewsletters []*dbmodel.NewsletterInfo
	err := pgxscan.Select(
		ctx,
		r.pool,
		&dbNewsletters,
		query.ReadNewsletterInfoWithLimit,
		lim,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query newsletters: %w", err)
	}

	// Překlad db modelů na servisní modely
	var svcNewsletters []*model.NewsletterInfo
	for _, dbNl := range dbNewsletters {
		var description *string // Inicializujte pointer na string
		if dbNl.Description.Valid {
			description = &dbNl.Description.String // Přiřaďte adresu, pokud je hodnota Valid
		}

		svcNl := &model.NewsletterInfo{
			Id:          dbNl.Id,
			Name:        dbNl.Name,
			Description: description, // Použijte pointer na string
			EditorId:    dbNl.EditorId,
			EditorEmail: dbNl.EditorEmail,
		}
		svcNewsletters = append(svcNewsletters, svcNl)
	}

	return svcNewsletters, nil
}

func (r *NewsletterRepository) UpdateNewsletter(ctx context.Context, newsletterId id.ID, name string, description string) (*model.BaseNewsletter, error) {
	var newsletter dbmodel.Newsletter

	namedArgs := pgx.NamedArgs{
		"id":   newsletterId,
		"name": name,
	}

	if description != "" {
		namedArgs["description"] = description
	} else {
		namedArgs["description"] = nil
	}
	if err := r.pool.QueryRow(
		ctx,
		mutation.UpdateNewsletter,
		namedArgs,
	).Scan(&newsletter.Id, &newsletter.Name, &newsletter.Description); err != nil {
		return nil, errors.New("error updating newsletter in DB: " + err.Error())
	}

	var nullableDescription *string
	if newsletter.Description.Valid {
		nullableDescription = &newsletter.Description.String
	}

	return &model.BaseNewsletter{
		ID:          newsletter.Id,
		Name:        newsletter.Name,
		Description: nullableDescription,
	}, nil
}
