package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"newsletterProject/config"
	"newsletterProject/transport/api/v1/model"
)

var cfg = config.MustLoadConfig()

func PostSupabaseOTPRequest(payload *model.SupabaseOTPPayload) (int, error) {
	jsonBytes, err := json.Marshal(payload)
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

	return 200, nil
}

func PostSupabaseRefreshRequest(token string) (*model.SupabaseRefreshResponse, int, error) {
	body := model.SupabaseRefreshPayload{
		RefreshToken: token,
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, 0, errors.New("Error marshalling JSON: " + err.Error())
	}

	request, err := http.NewRequest("POST", cfg.SupabaseURL+"/auth/v1/token?grant_type=refresh_token", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, 0, errors.New("Error creating request: " + err.Error())
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("apiKey", cfg.SupabaseAPIKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, 0, errors.New("HTTP request failed: " + err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		responseBody, readErr := io.ReadAll(response.Body)
		if readErr != nil {
			return nil, response.StatusCode, errors.New("Failed to read response body: " + readErr.Error())
		}
		return nil, response.StatusCode, fmt.Errorf("failed to post refresh request to Supabase: " + string(responseBody))
	}

	var responseBody model.SupabaseRefreshResponse
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return nil, 0, errors.New("Error decoding JSON response: " + err.Error())
	}

	return &responseBody, 200, nil
}
