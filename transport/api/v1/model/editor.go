package model

type Editor struct {
	Email string `json:"email" validate:"required,email"`
}
