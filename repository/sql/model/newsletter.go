package model

import (
	"database/sql"
	"newsletterProject/pkg/id"
)

type Newsletter struct {
	Id          id.ID          `db:"id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	EditorId    id.ID          `db:"editor_id"`
}

type NewsletterInfo struct {
	Id          id.ID          `db:"id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	EditorId    id.ID          `db:"editor_id"`
	EditorEmail string         `db:"editor_email"`
}
