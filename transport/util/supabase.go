package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"newsletterProject/config"
	"newsletterProject/transport/api/v1/model"
)

func PostSupabaseOTPRequest(email string) (int, error) {
	body := model.SupabasePayload{
		Email: email,
	}
	cfg := config.MustLoadConfig()

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return 0, errors.New("Error marshalling JSON: " + err.Error())
	}

	request, err := http.NewRequest("POST", cfg.SupabaseURL, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return 0, errors.New("Error creating request: " + err.Error())
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("apiKey", cfg.SupabaseAPIKey)

	client := &http.Client{}
	responsePayload, err := client.Do(request)
	if responsePayload.StatusCode < 200 || responsePayload.StatusCode >= 300 {
		body, readErr := io.ReadAll(responsePayload.Body)
		if readErr != nil {
			return responsePayload.StatusCode, errors.New("failed to post OTP request to supabase and failed to read the response body")
		}
		return responsePayload.StatusCode, errors.New("failed to post OTP request to supabase: " + string(body))
	}

	return 0, nil
}
