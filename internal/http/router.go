package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *Handler) http.Handler {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.CreateUser)
		r.Get("/{id}", h.GetUser)
	})

	return r
}
