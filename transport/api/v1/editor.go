package v1

import (
	"encoding/json"
	"net/http"
	apiEditor "newsletterProject/transport/api/v1/model"
	"newsletterProject/transport/util"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var transportEditor apiEditor.Editor
	err := json.NewDecoder(r.Body).Decode(&transportEditor)
	if err != nil {
		util.WriteResponse(w, http.StatusBadRequest, "Corrupted data")
	}

	editor, err := h.service.GetEditorByEmail(r.Context(), transportEditor.Email)
	if err != nil {
		util.WriteResponse(w, http.StatusNotFound, "Editor not found")
		return
	}

	// TODO: authorization with JWT token

	// testing
	if editor.Password == transportEditor.Password {
		util.WriteResponse(w, http.StatusOK, "Correct password")
	} else {
		util.WriteResponse(w, http.StatusNotFound, "Incorrect password")
	}
	return
}
