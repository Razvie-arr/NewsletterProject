package v1

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
	"newsletterProject/transport/api/v1/model"
	"newsletterProject/transport/util"
)

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		util.WriteResponse(w, http.StatusBadRequest, err)
		return
	}
	// Validate the post
	if err := validator.New().Struct(&post); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
}
