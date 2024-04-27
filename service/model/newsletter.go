package model

import "newsletterProject/pkg/id"

type Newsletter struct {
	ID          id.ID
	Name        string
	Description *string
	Editor      Editor
	Subscribers []Subscriber
}

type BaseNewsletter struct {
	ID          id.ID
	Name        string
	Description *string
}

type NewsletterInfo struct {
	Id          id.ID
	Name        string
	Description *string
	EditorId    id.ID
	EditorEmail string
}
