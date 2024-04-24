package v1

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"newsletterProject/mailer"
	"newsletterProject/pkg/id"
	"newsletterProject/service/errors"
	"newsletterProject/transport/api/v1/model"
	"newsletterProject/transport/util"
)

const createdSuccessfullyButNoEmailSent = "Post created successfully, but email was not sent"

func (h *Handler) PublishPost(w http.ResponseWriter, r *http.Request) {
	var apiPost model.Post
	if err := json.NewDecoder(r.Body).Decode(&apiPost); err != nil {
		util.WriteResponse(w, http.StatusBadRequest, err)
		return
	}
	// Validate the post
	if err := validator.New().Struct(&apiPost); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	// Check that newsletter exists
	newsletterUUID, err := uuid.Parse(apiPost.NewsletterId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, errors.ErrNewsletterNotFound)
		return
	}
	newsletterSvc, err := h.service.GetNewsletterById(r.Context(), id.ID(newsletterUUID))
	if err != nil {
		util.WriteErrResponse(w, http.StatusNotFound, err)
		return
	}

	// Check if current editor is newsletter's owner
	if newsletterSvc.Editor.ID.String() != r.Context().Value("editor_id") {
		util.WriteErrResponse(w, http.StatusForbidden, errors.ErrEditorIsNotOwner)
		return
	}

	post, err := h.service.PublishPost(r.Context(), apiPost.Content, newsletterSvc.ID)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	// send post via email to all subscribers
	var subscribersEmails = make([]string, len(newsletterSvc.Subscribers))
	for _, subscriber := range newsletterSvc.Subscribers {
		subscribersEmails = append(subscribersEmails, subscriber.Email)
	}
	subject := "New post in " + newsletterSvc.Name
	body, err := mailer.GetNewPostBody(newsletterSvc, post, "asd")
	if err != nil {
		// Post created successfully, but email was not sent
		util.WriteResponse(w, http.StatusOK, createdSuccessfullyButNoEmailSent)
		return
	}
	if err = h.service.SendEmail(subscribersEmails, subject, body); err != nil {
		util.WriteResponse(w, http.StatusOK, createdSuccessfullyButNoEmailSent)
		return
	}

}
