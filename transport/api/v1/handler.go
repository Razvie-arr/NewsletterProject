package v1

import (
	"github.com/go-chi/chi"
	"newsletterProject/transport/middleware"
)

type Handler struct {
	*chi.Mux

	authenticator middleware.Authenticator
	service       Service
}

func NewHandler(
	authenticator middleware.Authenticator,
	service Service,
) *Handler {
	h := &Handler{
		authenticator: authenticator,
		service:       service,
	}
	h.initRouter()
	return h
}

func (h *Handler) initRouter() {
	r := chi.NewRouter()

	authenticate := middleware.NewAuthenticate(h.authenticator)

	r.Route("/editor", func(r chi.Router) {
		r.Post("/register", h.Register)
		r.Post("/login", h.Login)
	})

	r.Route("/test", func(r chi.Router) {
		r.With(authenticate).Get("/", h.Test)
	})

	r.Route("/subscription", func(r chi.Router) {
		r.Post("/{newsletterId}", h.Subscribe)
		r.Delete("/{newsletterId}", h.Unsubscribe)
	})

	r.Route("/newsletter", func(r chi.Router) {
		r.Get("/unsubscribe", h.UnsubscribePage)
	})

	h.Mux = r
}
