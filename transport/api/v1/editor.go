package v1

import (
	"encoding/json"
	"fmt"
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
	if err != nil {
		util.WriteResponse(w, http.StatusNotFound, "Editor not found")
		return
	}

	// testing
	if editor.Password == transportEditor.Password {
		util.WriteResponse(w, http.StatusOK, "Correct password")
	} else {
		util.WriteResponse(w, http.StatusForbidden, "Incorrect password")
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
