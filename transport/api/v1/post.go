package v1

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"newsletterProject/mailer"
	"newsletterProject/pkg/id"
	"newsletterProject/service/errors"
	svcmodel "newsletterProject/service/model"
	"newsletterProject/transport/api/v1/model"
	"newsletterProject/transport/util"
	"strings"
)

const (
	createdSuccessfullyButNoEmailSent = "Post created successfully, but email was not sent to all users, not sent users: "
	createdSentSuccessfully           = "Post sent successfully"
)

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
	subject := "New post in " + newsletterSvc.Name
	var notSentUsers []string

	// Create a channel to collect errors from goroutines
	errCh := make(chan string, len(newsletterSvc.Subscribers))

	for _, subscriber := range newsletterSvc.Subscribers {
		go func(subscriber svcmodel.Subscriber) {
			email := subscriber.Email
			verificationString, err := h.service.GetVerificationString(r.Context(), newsletterSvc.ID, subscriber.ID)
			if err != nil {
				errCh <- email
				return
			}

			unsubscribeLink := mailer.GetUnsubscribeLink(newsletterUUID.String(), email, verificationString)
			body, err := mailer.GetNewPostBody(newsletterSvc, post, unsubscribeLink)
			if err != nil {
				errCh <- email
				return
			}
			to := []string{subscriber.Email}
			if err = h.service.SendEmail(to, subject, body); err != nil {
				errCh <- email
			} else {
				errCh <- "" // Send an empty string to the channel if there's no error
			}
		}(subscriber)
	}

	// Collect errors from the error channel
	for range newsletterSvc.Subscribers {
		email := <-errCh
		if email != "" {
			notSentUsers = append(notSentUsers, email)
		}
	}

	if len(notSentUsers) != 0 {
		notSentUsersString := strings.Join(notSentUsers, ", ")
		util.WriteResponse(w, http.StatusInternalServerError, createdSuccessfullyButNoEmailSent+notSentUsersString)
		return
	}
	util.WriteResponse(w, http.StatusOK, createdSentSuccessfully)
}
