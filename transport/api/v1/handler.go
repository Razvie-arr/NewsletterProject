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
		r.Get("/verify", h.Verify)
		r.Post("/login", h.Login)
		r.With(authenticate).Post("/refresh", h.Refresh)
	})

	r.Route("/subscription", func(r chi.Router) {
		r.Post("/{newsletterId}", h.Subscribe)
		r.Delete("/{newsletterId}", h.Unsubscribe)
	})

	r.Route("/newsletter", func(r chi.Router) {
		r.With(authenticate).Post("/", h.CreateNewsletter)
		r.With(authenticate).Delete("/{newsletterId}", h.DeleteNewsletter)
		r.With(authenticate).Patch("/", h.UpdateNewsletter)
		r.Get("/{newsletterId}", h.GetNewsletterById)
		r.Get("/", h.GetNewsletters)
		r.Get("/unsubscribe", h.UnsubscribePage)
	})

	r.Route("/post", func(r chi.Router) {
		r.With(authenticate).Post("/", h.PublishPost)
	})
	h.Mux = r
}
