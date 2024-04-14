package model

import "newsletterProject/pkg/id"

type Editor struct {
	Id       id.ID  `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
