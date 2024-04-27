package v1

import (
	"context"
	"newsletterProject/pkg/id"
	svcmodel "newsletterProject/service/model"
)

type Service interface {
	GetEditor(ctx context.Context, editorId id.ID) (*svcmodel.Editor, error)
	GetEditorByEmail(ctx context.Context, email string) (*svcmodel.Editor, error)
	CreateEditor(ctx context.Context, uuid id.ID, email string) (*svcmodel.Editor, error)
	Subscribe(ctx context.Context, newsletterId id.ID, subscriberMail string) (*svcmodel.Newsletter, string, error)
	SendEmail(to []string, subject, body string) error
	GetNewsletterById(ctx context.Context, newsletterId id.ID) (*svcmodel.Newsletter, error)
	GetVerificationString(ctx context.Context, newsletterId id.ID, subscriberId id.ID) (string, error)
	Unsubscribe(ctx context.Context, newsletterId id.ID, subscriberId id.ID) error
	PublishPost(ctx context.Context, content string, newsletterId id.ID) (*svcmodel.Post, error)
	CreateNewsletter(ctx context.Context, name, description string, editorId id.ID) (*svcmodel.BaseNewsletter, error)
	ExistsNewsletterOwnedByEditor(ctx context.Context, newsletterId id.ID, editorId id.ID) error
	DeleteNewsletter(ctx context.Context, newsletterId id.ID) error
	GetNewslettersInfo(ctx context.Context, limit int) ([]*svcmodel.NewsletterInfo, error)
	UpdateNewsletter(ctx context.Context, newsletterId id.ID, name, description string) (*svcmodel.BaseNewsletter, error)
}
