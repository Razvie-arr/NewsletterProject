package model

type Post struct {
	Content      string `json:"content" validate:"required"`
	NewsletterId string `json:"newsletter_id" validate:"required"`
}
