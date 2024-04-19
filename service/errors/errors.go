package errors

import "errors"

var (
	ErrAlreadySubscribed  = errors.New("email already subscribed")
	ErrNewsletterNotFound = errors.New("newsletter not found")
)
