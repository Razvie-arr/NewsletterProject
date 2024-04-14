package v1

import (
	"context"
	"newsletterProject/pkg/id"
	svcmodel "newsletterProject/service/model"
)

type Service interface {
	GetEditor(ctx context.Context, editorId id.ID) (*svcmodel.Editor, error)
}
