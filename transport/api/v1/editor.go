package v1

import (
	"encoding/json"
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
		util.WriteResponse(w, http.StatusNotFound, "Incorrect password")
	}
	return
}
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	//TODO: Send OTP to supabase
}

func (h *Handler) Test(w http.ResponseWriter, _ *http.Request) {
	util.WriteResponse(w, http.StatusAccepted, "This seems to work...")
	return
}
