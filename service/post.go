package service

import (
	"context"
	"newsletterProject/pkg/id"
	"newsletterProject/service/model"
)

func (s Service) PublishPost(ctx context.Context, content string, newsletterId id.ID) (*model.Post, error) {
	post, err := s.repository.CreatePost(ctx, content, newsletterId)
	if err != nil {
		return nil, err
	}

	return post, nil
}
