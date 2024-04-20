package service

import (
	"context"
	"newsletterProject/pkg/id"
	"newsletterProject/service/errors"
	"newsletterProject/service/model"
)

func (s Service) CreatePost(ctx context.Context, content string, editorId, newsletterId id.ID) (*model.Post, error) {
	// check if newsletter exists
	newsletter, err := s.repository.ReadNewsletter(ctx, newsletterId)
	if err != nil {
		return nil, errors.ErrNewsletterNotFound
	}

	if editorId != newsletter.Editor.ID {
		return nil, errors.ErrEditorIsNotOwner
	}

	post, err := s.repository.CreatePost(ctx, content, newsletterId)
	if err != nil {
		return nil, err
	}

	return post, nil
}
