package model

import "newsletterProject/pkg/id"

type Newsletter struct {
	ID          id.ID
	Name        string
	Description *string
	Editor      Editor
	Subscriber  []Subscriber
}
