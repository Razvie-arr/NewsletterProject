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

	// here should be routes

	h.Mux = r
}
