package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"newsletterProject/mailer"
	"newsletterProject/pkg/id"
	transportModel "newsletterProject/transport/api/v1/model"
	"newsletterProject/transport/util"
	"strings"
)

type verifyData struct {
	Email string `validate:"required,email"`
	Uuid  id.ID  `validate:"required"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var transportEditor transportModel.Editor
	err := json.NewDecoder(r.Body).Decode(&transportEditor)
	if err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Error reading request body: "+err.Error())
		return
	}

	// Validate the payload
	if err := validator.New().Struct(&transportEditor); err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Error validating request body: "+err.Error())
		return
	}

	editor, err := h.service.GetEditorByEmail(r.Context(), transportEditor.Email)
	if err != nil {
		util.WriteResponse(w, http.StatusNotFound, "Editor not found")
		return
	}

	payload := transportModel.SupabaseOTPPayload{
		Email: editor.Email,
	}

	// Validate the payload
	if err := validator.New().Struct(&payload); err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Error validating payload body: "+err.Error())
		return
	}

	if statusCode, err := requestOTP(&payload); err != nil {
		util.WriteErrResponse(w, statusCode, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, "login successful")
	return
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var payload transportModel.SupabaseOTPPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Error reading request body: "+err.Error())
		return
	}

	// Validate the payload
	if err := validator.New().Struct(&payload); err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Error validating request body: "+err.Error())
		return
	}

	//Check that email is not already in use
	if _, err := h.service.GetEditorByEmail(r.Context(), payload.Email); err == nil {
		util.WriteResponse(w, http.StatusConflict, "Email already in use")
		return
	}

	if statusCode, err := requestOTP(&payload); err != nil {
		util.WriteErrResponse(w, statusCode, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, "registration request successful")
	return
}

func requestOTP(payload *transportModel.SupabaseOTPPayload) (int, error) {
	statusCode, err := util.PostSupabaseOTPRequest(payload)
	if statusCode != 200 && err != nil {
		return statusCode, errors.New("Error sending request to supabase: " + err.Error())
	}
	if err != nil {
		return http.StatusInternalServerError, errors.New("Error sending request to supabase: " + err.Error())
	}
	if statusCode != 200 {
		return http.StatusInternalServerError, errors.New(fmt.Sprintf("Something unexpected happened: error is nil but statusCode is %d", statusCode))
	}
	return 200, nil
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

	http.Redirect(w, r, "/api/v1/editor/showJWT", http.StatusFound)
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

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	var payload transportModel.SupabaseRefreshPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Error reading request body: "+err.Error())
		return
	}

	// Validate the payload
	if err := validator.New().Struct(&payload); err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Error validating request body: "+err.Error())
		return
	}

	requestSessionRefresh(w, payload.RefreshToken)
	return
}

func requestSessionRefresh(w http.ResponseWriter, token string) {
	response, statusCode, err := util.PostSupabaseRefreshRequest(token)
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

	if err := validator.New().Struct(response); err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Error validating response from supabase: "+err.Error())
		return
	}

	util.WriteResponseWithJsonBody(w, http.StatusOK, response)

	return
}

func (h *Handler) ShowJWTPage(w http.ResponseWriter, _ *http.Request) {
	page, err := mailer.GetShowJWTPageBody()
	if err != nil {
		util.WriteResponse(w, http.StatusInternalServerError, "Error generating page: "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(page))
}
