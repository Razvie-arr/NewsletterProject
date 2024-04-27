package model

import "newsletterProject/pkg/id"

type Newsletter struct {
	Id          id.ID  `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type NewsletterInfo struct {
	Id          id.ID   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	EditorId    id.ID   `json:"editorId"`
	EditorEmail string  `json:"editorEmail"`
}
