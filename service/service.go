package service

import (
	"context"
	"newsletterProject/mailer"
	"newsletterProject/pkg/id"
	"newsletterProject/service/model"
)

type Repository interface {
	// Editor
	ReadEditor(ctx context.Context, editorId id.ID) (*model.Editor, error)
	ReadEditorByEmail(ctx context.Context, email string) (*model.Editor, error)
	// Subscriber
	ReadSubscriberByEmail(ctx context.Context, email string) (*model.Subscriber, error)
	CreateSubscriber(ctx context.Context, email string) (*model.Subscriber, error)
	// Newsletter
	ReadNewsletter(ctx context.Context, newsletterId id.ID) (*model.Newsletter, error)
	Subscribe(ctx context.Context, newsletterId id.ID, subscriberId id.ID) (string, error)
	GetVerificationString(ctx context.Context, newsletterId id.ID, subscriberId id.ID) (string, error)
	Unsubscribe(ctx context.Context, newsletterId id.ID, subscriberId id.ID) error
}

type Service struct {
	repository Repository
	mailer     mailer.Mailer
}

func NewService(
	repository Repository,
	mailer mailer.Mailer,
) (Service, error) {
	return Service{
		repository: repository,
		mailer:     mailer,
	}, nil
}

func (s Service) SendEmail(to []string, subject, body string) error {
	return s.mailer.SendEmail(to, subject, body)
}
