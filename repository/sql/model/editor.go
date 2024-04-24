package model

import "newsletterProject/pkg/id"

type Editor struct {
	Id    id.ID  `db:"uuid"`
	Email string `db:"email"`
}
