package model

type SupabaseOTPPayload struct {
	Email string `json:"email" validate:"required,email"`
}

type SupabaseRefreshPayload struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type SupabaseRefreshResponse struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}
