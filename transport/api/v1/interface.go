package v1

import (
	"context"
	"newsletterProject/pkg/id"
	svcmodel "newsletterProject/service/model"
)

type Service interface {
	GetEditor(ctx context.Context, editorId id.ID) (*svcmodel.Editor, error)
	GetEditorByEmail(ctx context.Context, email string) (*svcmodel.Editor, error)
	CreateEditor(ctx context.Context, email, password string) (*svcmodel.Editor, error)
	Subscribe(ctx context.Context, newsletterId id.ID, subscriberMail string) (*svcmodel.Newsletter, string, error)
	SendEmail(to []string, subject, body string) error
	GetNewsletterById(ctx context.Context, newsletterId id.ID) (*svcmodel.Newsletter, error)
	GetVerificationString(ctx context.Context, newsletterId id.ID, subscriberId id.ID) (string, error)
	Unsubscribe(ctx context.Context, newsletterId id.ID, subscriberId id.ID) error
}
