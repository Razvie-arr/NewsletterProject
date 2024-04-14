package v1

import "github.com/go-chi/chi"

type Handler struct {
	*chi.Mux

	service Service
}

func NewHandler(
	service Service,
) *Handler {
	h := &Handler{
		service: service,
	}
	h.initRouter()
	return h
}

func (h *Handler) initRouter() {
	r := chi.NewRouter()

	// TODO: Setup middleware.

	r.Route("/editor", func(r chi.Router) {
		r.Post("/login", h.Login)
	})
	h.Mux = r
}
