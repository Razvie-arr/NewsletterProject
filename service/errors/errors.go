package errors

import "errors"

var (
	ErrAlreadySubscribed  = errors.New("email already subscribed")
	ErrNewsletterNotFound = errors.New("newsletter not found")
	ErrEditorIsNotOwner   = errors.New("editor is not owner of newsletter")
)
