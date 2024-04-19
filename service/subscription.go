package service

import (
	"context"
	"newsletterProject/pkg/id"
	"newsletterProject/service/errors"
	svcmodel "newsletterProject/service/model"
)

func (s Service) Subscribe(ctx context.Context, newsletterId id.ID, subscriberMail string) (*svcmodel.Newsletter, string, error) {
	// Get the subscriber by email
	subscriber, err := s.repository.ReadSubscriberByEmail(ctx, subscriberMail)
	if err != nil {
		// Subscriber not found, we create a new one
		sub, err := s.repository.CreateSubscriber(ctx, subscriberMail)
		if err != nil {
			return nil, "", err
		}
		subscriber = sub
	}

	// Check if the newsletter exists
	newsletter, err := s.repository.ReadNewsletter(ctx, newsletterId)
	if err != nil {
		return nil, "", errors.ErrNewsletterNotFound
	}

	// Check if the subscriber is already subscribed to the newsletter
	for _, sub := range newsletter.Subscriber {
		if sub.ID == subscriber.ID {
			return nil, "", errors.ErrAlreadySubscribed
		}
	}

	// Subscribe the subscriber to the newsletter
	verificationString, err := s.repository.Subscribe(ctx, newsletterId, subscriber.ID)
	if err != nil {
		return nil, "", err
	}

	return newsletter, verificationString, nil
}

func (s Service) GetVerificationString(ctx context.Context, newsletterId id.ID, subscriberId id.ID) (string, error) {
	verificationString, err := s.repository.GetVerificationString(ctx, newsletterId, subscriberId)
	if err != nil {
		return "", err
	}
	return verificationString, nil
}

func (s Service) Unsubscribe(ctx context.Context, newsletterId id.ID, subscriberId id.ID) error {
	err := s.repository.Unsubscribe(ctx, newsletterId, subscriberId)
	if err != nil {
		return err
	}
	return nil
}
