package mailer

type SuccessfulEmailHTMLParams struct {
	NewsletterName        string
	NewsletterDescription string
	EditorMail            string
	UnsubscribeLink       string
}

type NewPostEmailHTMLParams struct {
	NewsletterName        string
	NewsletterDescription string
	EditorMail            string
	PostContent           string
	UnsubscribeLink       string
}

type UnsubscribePageHTMLParams struct {
	NewsletterId       string
	SubscriberEmail    string
	VerificationString string
}
