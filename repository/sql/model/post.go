package model

import (
	"newsletterProject/pkg/id"
)

type Post struct {
	Id           id.ID  `db:"id"`
	Content      string `db:"content"`
	NewsletterId id.ID  `db:"newsletter_id"`
}
