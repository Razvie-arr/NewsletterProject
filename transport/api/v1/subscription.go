package v1

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"newsletterProject/mailer"
	"newsletterProject/pkg/id"
	svcErrors "newsletterProject/service/errors"
	"newsletterProject/transport/api/v1/model"
	"newsletterProject/transport/util"
)

const (
	subscriptionConfirmation          = "Subscription confirmation"
	subscriptionSuccessfulNoEmailSent = "Subscription successful, but email was not sent"
	subscriptionSuccessful            = "Subscription successful"
)

func (h *Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	newsletterId := chi.URLParam(r, "newsletterId")
	var Subscriber model.Subscriber
	if err := json.NewDecoder(r.Body).Decode(&Subscriber); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	// Validate the subscriber
	if err := validator.New().Struct(&Subscriber); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	newsletterUUID, err := uuid.Parse(newsletterId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, errors.New("newsletter not found"))
		return
	}

	newsletter, verificationString, err := h.service.Subscribe(r.Context(), id.ID(newsletterUUID), Subscriber.Email)
	if err != nil {
		switch {
		case errors.Is(err, svcErrors.ErrAlreadySubscribed):
			util.WriteErrResponse(w, http.StatusConflict, err)
		case errors.Is(err, svcErrors.ErrNewsletterNotFound):
			util.WriteErrResponse(w, http.StatusNotFound, err)
		default:
			util.WriteErrResponse(w, http.StatusInternalServerError, err)
		}
		return
	}

	// Try to send confirmation email
	to := []string{Subscriber.Email}
	subject := subscriptionConfirmation
	unsubLink := mailer.GetUnsubscribeLink(newsletterId, Subscriber.Email, verificationString)
	body, err := mailer.GetSuccessfulSubscriptionEmailBody(newsletter, unsubLink)
	if err != nil {
		// Subscription was successful, but email was not sent
		util.WriteResponse(w, http.StatusOK, subscriptionSuccessfulNoEmailSent)
	}
	if err := h.service.SendEmail(to, subject, body); err != nil {
		// Subscription was successful, but email was not sent
		util.WriteResponse(w, http.StatusOK, subscriptionSuccessfulNoEmailSent)
	}

	util.WriteResponse(w, http.StatusOK, subscriptionSuccessful)
}

func (h *Handler) UnsubscribePage(w http.ResponseWriter, r *http.Request) {
	newsletterId := r.URL.Query().Get("newsletterId")
	email := r.URL.Query().Get("email")
	verificationString := r.URL.Query().Get("verificationString")

	page, err := mailer.GetUnsubscribePageBody(newsletterId, email, verificationString)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(page))
}

func (h *Handler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	newsletterId := chi.URLParam(r, "newsletterId")
	var body model.UnsubscribeRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	// Validate the body
	if err := validator.New().Struct(&body); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	// Check that newsletter exists
	newsletterUUID, err := uuid.Parse(newsletterId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, errors.New("newsletter not found"))
		return
	}

	newsletter, err := h.service.GetNewsletterById(r.Context(), id.ID(newsletterUUID))
	if err != nil {
		util.WriteErrResponse(w, http.StatusNotFound, err)
		return
	}

	// Check if subscriber is subscribed to the newsletter
	for _, subscriber := range newsletter.Subscriber {
		if subscriber.Email == body.Email {
			// Verify string
			dbVerificationString, err := h.service.GetVerificationString(r.Context(), id.ID(newsletterUUID), subscriber.ID)
			if err != nil {
				util.WriteErrResponse(w, http.StatusInternalServerError, err)
				return
			}
			if dbVerificationString == body.VerificationString {
				// Unsubscribe
				if err := h.service.Unsubscribe(r.Context(), id.ID(newsletterUUID), subscriber.ID); err != nil {
					util.WriteErrResponse(w, http.StatusInternalServerError, err)
					return
				}
				util.WriteResponse(w, http.StatusOK, "Unsubscribed successfully")
				return
			}
			// Verification string does not match
			util.WriteErrResponse(w, http.StatusUnauthorized, errors.New("verification string does not match"))
			return
		}
	}
	// No subscription found
	util.WriteErrResponse(w, http.StatusNotFound, errors.New("no subscription found for this email"))
}
