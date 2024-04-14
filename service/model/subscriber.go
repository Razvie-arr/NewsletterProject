package model

import "newsletterProject/pkg/id"

type Subscriber struct {
	ID         id.ID
	Email      string
	Newsletter []Newsletter
}
