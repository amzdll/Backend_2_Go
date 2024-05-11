package client

import (
	"github.com/go-chi/chi/v5"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/clients", h.CreateClient)
	r.Get("/clients", h.GetAllClients)

	return r
}
