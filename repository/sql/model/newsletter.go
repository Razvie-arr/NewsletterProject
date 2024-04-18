package model

import "newsletterProject/pkg/id"

type Newsletter struct {
	Id          id.ID  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	EditorId    id.ID  `db:"editor_id"`
}
