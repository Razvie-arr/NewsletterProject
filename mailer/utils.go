package mailer

import (
	"bytes"
	"html/template"
	"newsletterProject/service/model"
	"os"
	"path/filepath"
)

func GetUnsubscribePageBody(newsletterId, subscriberMail, verificationString string) (string, error) {
	data := UnsubscribePageHTMLParams{
		NewsletterId:       newsletterId,
		SubscriberEmail:    subscriberMail,
		VerificationString: verificationString,
	}
	cwd, _ := os.Getwd()
	templatePath := filepath.Join(cwd, "mailer/templates/unsubscribePage.html")
	t, _ := template.ParseFiles(templatePath)
	var page bytes.Buffer
	if err := t.Execute(&page, data); err != nil {
		return "", err
	}
	return page.String(), nil
}

func GetSuccessfulSubscriptionEmailBody(newsletter *model.Newsletter, subscriberEmail, verificationString string) (string, error) {
	var description string
	if newsletter.Description != nil {
		description = *newsletter.Description
	}
	data := SuccessfulEmailHTMLParams{
		NewsletterName:        newsletter.Name,
		NewsletterDescription: description,
		NewsletterId:          newsletter.ID.String(),
		EditorMail:            newsletter.Editor.Email,
		SubscriberEmail:       subscriberEmail,
		VerificationString:    verificationString,
	}
	cwd, _ := os.Getwd()
	templatePath := filepath.Join(cwd, "mailer/templates/subscribeConfirmation.html")
	t, _ := template.ParseFiles(templatePath)
	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return "", err
	}

	return body.String(), nil
}

func GetNewPostBody(newsletter *model.Newsletter, post *model.Post, subscriberEmail, verificationString string) (string, error) {
	var description string
	if newsletter.Description != nil {
		description = *newsletter.Description
	}
	data := NewPostEmailHTMLParams{
		NewsletterId:          newsletter.ID.String(),
		NewsletterName:        newsletter.Name,
		NewsletterDescription: description,
		EditorMail:            newsletter.Editor.Email,
		PostContent:           post.Text,
		SubscriberEmail:       subscriberEmail,
		VerificationString:    verificationString,
	}
	cwd, _ := os.Getwd()
	templatePath := filepath.Join(cwd, "mailer/templates/newPost.html")
	t, _ := template.ParseFiles(templatePath)
	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return "", err
	}
	return body.String(), nil
}
