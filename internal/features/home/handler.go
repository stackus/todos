package home

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/stackus/todos/internal/templates/pages"
)

type (
	Handler interface {
		// Home : GET /
		Home(w http.ResponseWriter, r *http.Request)
	}

	handler struct {
		service Service
	}
)

func NewHandler(svc Service) Handler {
	return &handler{service: svc}
}

func Mount(r chi.Router, h Handler) {
	r.Get("/", h.Home)
}

func (h handler) Home(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := pages.HomePage(todos).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
