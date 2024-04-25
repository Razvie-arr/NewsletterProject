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
