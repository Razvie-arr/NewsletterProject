package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"newsletterProject/pkg/id"
	apiEditor "newsletterProject/transport/api/v1/model"
	"newsletterProject/transport/util"
	"strings"
)

type verifyData struct {
	Email string `validate:"required,email"`
	Uuid  id.ID  `validate:"required"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var transportEditor apiEditor.Editor
	err := json.NewDecoder(r.Body).Decode(&transportEditor)
	if err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Corrupted data")
		return
	}

	editor, err := h.service.GetEditorByEmail(r.Context(), transportEditor.Email)
	if err != nil {
		util.WriteResponse(w, http.StatusNotFound, "Editor not found")
		return
	}

	requestOTP(w, editor.Email)

	util.WriteResponse(w, http.StatusOK, "login successful")
	return
}
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var payload apiEditor.SupabasePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Error reading request body")
		return
	}

	requestOTP(w, payload.Email)

	util.WriteResponse(w, http.StatusOK, "registration successful")
	return
}

func requestOTP(w http.ResponseWriter, email string) {
	statusCode, err := util.PostSupabaseOTPRequest(email)
	if statusCode != 200 && err != nil {
		util.WriteResponse(w, statusCode, "Error sending request to supabase: "+err.Error())
		return
	}
	if err != nil {
		util.WriteResponse(w, http.StatusInternalServerError, "Error sending request to supabase: "+err.Error())
		return
	}
	if statusCode != 200 {
		util.WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("Something unexpected happened: error is nil but statusCode is %d", statusCode))
		return
	}
	return
}

func (h *Handler) Verify(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	if data == "" {
		util.WriteResponse(w, http.StatusBadRequest, "No data provided")
		return
	}

	params, err := parseData(data)
	if err != nil {
		util.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if _, err := h.service.CreateEditor(r.Context(), params.Uuid, params.Email); err != nil {
		util.WriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.WriteResponse(w, http.StatusOK, "verification successful")
	return
}

func parseData(data string) (*verifyData, error) {

	// Remove the 'map[' prefix and the closing ']'
	trimmedData := strings.TrimPrefix(data, "map[")
	trimmedData = strings.TrimSuffix(trimmedData, "]")

	// Split the string into key-value pairs
	dataPairs := strings.Split(trimmedData, " ")

	// Struct to store and validate the values
	params := verifyData{}
	for _, pair := range dataPairs {
		// Split the pair by ':' to get key and value
		kv := strings.Split(pair, ":")
		if len(kv) == 2 {
			key := strings.TrimSpace(kv[0])
			value, err := url.QueryUnescape(kv[1])
			if err != nil {
				return nil, errors.New("failed to parse query parameters")
			}
			switch key {
			case "email":
				params.Email = value
			case "sub":
				userId, err := uuid.Parse(value)
				if err != nil {
					return nil, errors.New("failed to parse uuid: " + value)
				}
				params.Uuid = id.ID(userId)
			}
		}
	}

	// Validate struct
	err := validator.New().Struct(&params)
	if err != nil {
		return nil, errors.New("Invalid parameters provided: " + err.Error())
	}

	return &params, nil
}

func (h *Handler) Test(w http.ResponseWriter, _ *http.Request) {
	util.WriteResponse(w, http.StatusAccepted, "This seems to work...")
	return
}
