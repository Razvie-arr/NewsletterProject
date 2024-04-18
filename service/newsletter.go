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
