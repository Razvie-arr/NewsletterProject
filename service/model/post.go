package model

import "newsletterProject/pkg/id"

type Post struct {
	ID         id.ID
	Newsletter Newsletter
	Text       string
}
