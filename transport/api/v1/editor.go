package v1

import (
	"encoding/json"
	"net/http"
	apiEditor "newsletterProject/transport/api/v1/model"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var editor apiEditor.Editor
	err := json.NewDecoder(r.Body).Decode(&editor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// TODO: authorization with JWT token
}
