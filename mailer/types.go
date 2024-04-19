package mailer

type SuccessfulEmailHTMLParams struct {
	NewsletterId          string
	NewsletterName        string
	NewsletterDescription string
	EditorMail            string
	SubscriberEmail       string
	VerificationString    string
}

type NewPostEmailHTMLParams struct {
	NewsletterId          string
	NewsletterName        string
	NewsletterDescription string
	EditorMail            string
	PostContent           string
	SubscriberEmail       string
	VerificationString    string
}

type UnsubscribePageHTMLParams struct {
	NewsletterId       string
	SubscriberEmail    string
	VerificationString string
}
