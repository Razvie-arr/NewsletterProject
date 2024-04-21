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

	request, err := http.NewRequest("POST", cfg.SupabaseURL+"/auth/v1/otp", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return 0, errors.New("Error creating request: " + err.Error())
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("apiKey", cfg.SupabaseAPIKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return 0, errors.New("HTTP request failed: " + err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		// Read the response body only if the request was not successful
		responseBody, readErr := io.ReadAll(response.Body)
		if readErr != nil {
			return response.StatusCode, errors.New("failed to read response body: " + readErr.Error())
		}
		return response.StatusCode, errors.New("failed to post OTP request to Supabase: " + string(responseBody))
	}

	return 0, nil
}
