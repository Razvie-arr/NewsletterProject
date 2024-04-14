package service

import (
	"context"
	"newsletterProject/pkg/id"
	"newsletterProject/service/model"
)

func (s Service) GetEditor(ctx context.Context, editorId id.ID) (*model.Editor, error) {
	editor, err := s.repository.ReadEditor(ctx, editorId)
	if err != nil {
		return nil, err
	}
	return editor, nil
}

func (s Service) GetEditorByEmail(ctx context.Context, email string) (*model.Editor, error) {
	editor, err := s.repository.ReadEditorByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return editor, nil
}
