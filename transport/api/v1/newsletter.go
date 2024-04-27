package v1

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"newsletterProject/pkg/id"
	"newsletterProject/transport/api/v1/model"
	"newsletterProject/transport/util"
	"strconv"
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

func (h *Handler) DeleteNewsletter(w http.ResponseWriter, r *http.Request) {
	idUUID, err := uuid.Parse(chi.URLParam(r, "newsletterId"))
	// validation of newsletter ID
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, errors.New("invalid newsletter ID"))
		return
	}
	editorId := r.Context().Value("editor_id")
	editorIdUUID, err := uuid.Parse(editorId.(string))
	// validation of editor ID
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, errors.New("invalid credentials"))
		return
	}

	err = h.service.ExistsNewsletterOwnedByEditor(r.Context(), id.ID(idUUID), id.ID(editorIdUUID))
	if err != nil {
		util.WriteErrResponse(w, http.StatusNotFound, errors.New("newsletter owned by this editor was not found"))
		return
	}

	err = h.service.DeleteNewsletter(r.Context(), id.ID(idUUID))
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	util.WriteResponse(w, http.StatusNoContent, nil)
}

func (h *Handler) GetNewsletterById(w http.ResponseWriter, r *http.Request) {
	idUUID, err := uuid.Parse(chi.URLParam(r, "newsletterId"))
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, errors.New("invalid newsletter ID"))
		return

	}
	newsletter, err := h.service.GetNewsletterById(r.Context(), id.ID(idUUID))
	if err != nil {
		util.WriteErrResponse(w, http.StatusNotFound, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	util.WriteResponse(w, http.StatusOK, model.NewsletterInfo{
		Id:          newsletter.ID,
		Name:        newsletter.Name,
		Description: newsletter.Description,
		EditorId:    newsletter.Editor.ID,
		EditorEmail: newsletter.Editor.Email,
	})

}

func (h *Handler) GetNewsletters(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	limitInt := 50

	if limit != "" {
		var err error
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			util.WriteErrResponse(w, http.StatusBadRequest, errors.New("invalid limit"))
			return
		}
		if limitInt > 250 || limitInt < 1 {
			util.WriteErrResponse(w, http.StatusBadRequest, errors.New("limit must be between 1 and 250"))
			return
		}
	}
	svcNewsletters, err := h.service.GetNewslettersInfo(r.Context(), limitInt)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Překlad ze servisních modelů do transportních modelů
	var transNewsletters []model.NewsletterInfo
	for _, n := range svcNewsletters {
		transNewsletter := model.NewsletterInfo{
			Id:          n.Id,
			Name:        n.Name,
			Description: n.Description,
			EditorId:    n.EditorId,
			EditorEmail: n.EditorEmail,
		}
		transNewsletters = append(transNewsletters, transNewsletter)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(transNewsletters); err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, errors.New("failed to encode newsletters"))
		return
	}
}

func (h *Handler) UpdateNewsletter(w http.ResponseWriter, r *http.Request) {
	editorId := r.Context().Value("editor_id")
	editorIdUUID, err := uuid.Parse(editorId.(string))

	// validation of editor ID
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, errors.New("invalid credentials"))
		return
	}

	var newsletter model.Newsletter
	if err := json.NewDecoder(r.Body).Decode(&newsletter); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, errors.New("invalid body"))
		return
	}
	// validace
	if err := validator.New().Struct(&newsletter); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.ExistsNewsletterOwnedByEditor(r.Context(), newsletter.Id, id.ID(editorIdUUID))
	if err != nil {
		util.WriteErrResponse(w, http.StatusNotFound, errors.New("newsletter owned by this editor was not found"))
		return
	}

	updatedNewsletter, err := h.service.UpdateNewsletter(r.Context(), newsletter.Id, newsletter.Name, newsletter.Description)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	var description string
	if updatedNewsletter.Description != nil {
		description = *updatedNewsletter.Description
	}

	w.Header().Set("Content-Type", "application/json")
	util.WriteResponse(w, http.StatusOK, model.Newsletter{
		Id:          updatedNewsletter.ID,
		Name:        updatedNewsletter.Name,
		Description: description,
	})

}
