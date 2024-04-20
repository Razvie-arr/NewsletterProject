package model

type Post struct {
	Content      string `json:"content" validate:"required"`
	NewsletterId int    `json:"newsletter_id" validate:"required"`
}
