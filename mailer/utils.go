package mailer

import (
	"bytes"
	"html/template"
	"newsletterProject/service/model"
	"os"
	"path/filepath"
)

const (
	unsubscribePageTemplateName  = "unsubscribePage"
	subscribeConfirmTemplateName = "subscribeConfirmation"
	newPostTemplateName          = "newPost"
	showJWTPageTemplateName      = "showJWTPage"
)

func getTemplatePath(name string) string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "mailer/templates/"+name+".html")
}

func GetUnsubscribeLink(newsletterId, subscriberMail, verificationString string) string {
	return "http://16.16.58.152:" + os.Getenv("PORT") + "/api/v1/newsletter/unsubscribe?newsletterId=" + newsletterId + "&email=" + subscriberMail + "&verificationString=" + verificationString
}

func GetUnsubscribePageBody(newsletterId, subscriberMail, verificationString string) (string, error) {
	data := UnsubscribePageHTMLParams{
		NewsletterId:       newsletterId,
		SubscriberEmail:    subscriberMail,
		VerificationString: verificationString,
	}
	templatePath := getTemplatePath(unsubscribePageTemplateName)
	t, _ := template.ParseFiles(templatePath)
	var page bytes.Buffer
	if err := t.Execute(&page, data); err != nil {
		return "", err
	}
	return page.String(), nil
}

func GetSuccessfulSubscriptionEmailBody(newsletter *model.Newsletter, unsubscribeLink string) (string, error) {
	var description string
	if newsletter.Description != nil {
		description = *newsletter.Description
	}
	data := SuccessfulEmailHTMLParams{
		NewsletterName:        newsletter.Name,
		NewsletterDescription: description,
		EditorMail:            newsletter.Editor.Email,
		UnsubscribeLink:       unsubscribeLink,
	}
	templatePath := getTemplatePath(subscribeConfirmTemplateName)
	t, _ := template.ParseFiles(templatePath)
	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return "", err
	}

	return body.String(), nil
}

func GetNewPostBody(newsletter *model.Newsletter, post *model.Post, unsubscribeLink string) (string, error) {
	var description string
	if newsletter.Description != nil {
		description = *newsletter.Description
	}
	data := NewPostEmailHTMLParams{
		NewsletterName:        newsletter.Name,
		NewsletterDescription: description,
		EditorMail:            newsletter.Editor.Email,
		PostContent:           post.Content,
		UnsubscribeLink:       unsubscribeLink,
	}
	templatePath := getTemplatePath(newPostTemplateName)
	t, _ := template.ParseFiles(templatePath)
	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return "", err
	}
	return body.String(), nil
}

func GetShowJWTPageBody() (string, error) {
	templatePath := getTemplatePath(showJWTPageTemplateName)
	t, parseError := template.ParseFiles(templatePath)
	if parseError != nil {
		return "Error with template parsing with path" + templatePath, parseError
	}
	var body bytes.Buffer
	err := t.Execute(&body, nil)
	if err != nil {
		return "", err
	}
	return body.String(), nil
}
