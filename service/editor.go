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
