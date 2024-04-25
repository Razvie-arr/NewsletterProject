package v1

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"newsletterProject/pkg/id"
	"newsletterProject/transport/api/v1/model"
	"newsletterProject/transport/util"
)

func (h *Handler) CreateNewsletter(w http.ResponseWriter, r *http.Request) {
	editorId := r.Context().Value("editor_id")
	var newsletter model.Newsletter
	if err := json.NewDecoder(r.Body).Decode(&newsletter); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, errors.New("invalid body"))
		return
	}
	editorIdUUID, err := uuid.Parse(editorId.(string))
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, errors.New("invalid credentials"))
		return
	}
	newsletter.Id = id.ID(editorIdUUID)

	// Validate the newsletter
	if err := validator.New().Struct(&newsletter); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	createdNewsletter, err := h.service.CreateNewsletter(r.Context(), newsletter.Name, newsletter.Description, newsletter.Id)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	var description string
	if createdNewsletter.Description != nil {
		description = *createdNewsletter.Description
	}

	w.Header().Set("Content-Type", "application/json")

	util.WriteResponse(w, http.StatusCreated, model.Newsletter{
		Id:          createdNewsletter.ID,
		Name:        createdNewsletter.Name,
		Description: description,
	})
}
