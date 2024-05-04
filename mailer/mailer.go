package mailer

import "github.com/resend/resend-go/v2"

type Mailer interface {
	SendEmail(to []string, subject, body string) error
}

type ResendMailer struct {
	*resend.Client
}

func NewResendMailer(apiKey string) *ResendMailer {
	return &ResendMailer{
		Client: resend.NewClient(apiKey),
	}
}

func (r *ResendMailer) SendEmail(to []string, subject, body string) error {
	params := &resend.SendEmailRequest{
		// TODO: Use the registered domain here after we get it
		From:    "newsletter@resend.dev",
		To:      to,
		Subject: subject,
		Html:    body,
	}
	_, err := r.Client.Emails.Send(params)
	return err
}
