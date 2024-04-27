package service

import (
	"context"
	"newsletterProject/pkg/id"
	"newsletterProject/service/errors"
	svcmodel "newsletterProject/service/model"
)

func (s Service) GetNewsletterById(ctx context.Context, newsletterId id.ID) (*svcmodel.Newsletter, error) {
	newsletter, err := s.repository.ReadNewsletter(ctx, newsletterId)
	if err != nil {
		return nil, errors.ErrNewsletterNotFound
	}
	return newsletter, nil
}

func (s Service) CreateNewsletter(ctx context.Context, name, description string, editorId id.ID) (*svcmodel.BaseNewsletter, error) {
	svcNewsletter, err := s.repository.CreateNewsletter(ctx, name, description, editorId)
	if err != nil {
		return nil, err
	}
	return svcNewsletter, nil
}

func (s Service) DeleteNewsletter(ctx context.Context, newsletterId id.ID) error {
	return s.repository.DeleteNewsletter(ctx, newsletterId)
}

func (s Service) ExistsNewsletterOwnedByEditor(ctx context.Context, newsletterId id.ID, editorId id.ID) error {
	return s.repository.ExistsNewsletterWithEditor(ctx, newsletterId, editorId)
}

func (s Service) GetNewslettersInfo(ctx context.Context, limit int) ([]*svcmodel.NewsletterInfo, error) {
	svcNewsletterInfo, err := s.repository.GetNewslettersInfo(ctx, limit)
	if err != nil {
		return nil, err
	}
	return svcNewsletterInfo, nil
}

func (s Service) UpdateNewsletter(ctx context.Context, newsletterId id.ID, name, description string) (*svcmodel.BaseNewsletter, error) {
	svcNewsletter, err := s.repository.UpdateNewsletter(ctx, newsletterId, name, description)
	if err != nil {
		return nil, err
	}
	return svcNewsletter, nil
}
