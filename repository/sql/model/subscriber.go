package model

import "newsletterProject/pkg/id"

type Subscriber struct {
	Id    id.ID  `db:"id"`
	Email string `db:"email"`
}
