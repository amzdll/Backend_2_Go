package client

import (
	"github.com/go-chi/chi/v5"
)

type Service interface {
	ClientHello() string
}

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/clients", h.Create)
	r.Get("/clients/{name}_{surname}", h.GetByNameSurname)
	r.Get("/clients", h.GetAll)
	r.Put("/clients", h.UpdateAddress)
	r.Delete("/clients", h.DeleteById)

	return r
}
