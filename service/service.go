package service

import (
	"context"
	"newsletterProject/pkg/id"
	"newsletterProject/service/model"
)

type Repository interface {
	ReadEditor(ctx context.Context, editorId id.ID) (*model.Editor, error)
	ReadEditorByEmail(ctx context.Context, email string) (*model.Editor, error)
}

type Service struct {
	repository Repository
}

func NewService(
	repository Repository,
) (Service, error) {
	return Service{
		repository: repository,
	}, nil
}
