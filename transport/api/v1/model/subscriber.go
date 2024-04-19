package model

type Subscriber struct {
	Email string `json:"email" validate:"required,email"`
}

type UnsubscribeRequestBody struct {
	Email              string `json:"email" validate:"required,email"`
	VerificationString string `json:"verificationString" validate:"required"`
}
