package model

import "newsletterProject/pkg/id"

type Newsletter struct {
	Id          id.ID  `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
