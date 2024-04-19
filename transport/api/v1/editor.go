package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"net/http"
	apiEditor "newsletterProject/transport/api/v1/model"
	"newsletterProject/transport/util"
)

var transportEditor apiEditor.Editor

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&transportEditor)
	if err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Corrupted data")
	}

	editor, err := h.service.GetEditorByEmail(r.Context(), transportEditor.Email)
	if err == nil {
		// testing
		if editor.Password == transportEditor.Password {
			util.WriteResponse(w, http.StatusOK, "Correct password")
		} else {
			util.WriteResponse(w, http.StatusUnauthorized, "Incorrect password")
		}
		return
	}

	if !errors.Is(err, pgx.ErrNoRows) {
		util.WriteResponse(w, http.StatusInternalServerError, "Error reading DB: "+err.Error())
		return
	}

	if _, err := h.service.CreateEditor(r.Context(), transportEditor.Email, transportEditor.Password); err != nil {
		util.WriteResponse(w, http.StatusInternalServerError, "Error occurred while inserting to DB: "+err.Error())
		return
	}
	return
}
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var payload apiEditor.SupabasePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Error reading request body")
		return
	}

	statusCode, err := util.PostSupabaseOTPRequest(payload.Email)
	if statusCode != 0 && err != nil {
		util.WriteResponse(w, statusCode, "Error sending request to supabase: "+err.Error())
		return
	}
	if err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Error sending request to supabase: "+err.Error())
		return
	}
	if statusCode != 0 && err == nil {
		util.WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("Something unexpected happened: error is nil but statusCode is %d", statusCode))
		return
	}
	return
}

func (h *Handler) Test(w http.ResponseWriter, _ *http.Request) {
	util.WriteResponse(w, http.StatusAccepted, "This seems to work...")
	return
}
