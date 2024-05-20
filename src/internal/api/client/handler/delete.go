package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/nicklaw5/go-respond"
	"net/http"
)

func (h *Handler) DeleteById(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, err := uuid.Parse(idString)
	if err != nil {
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}
	if err = h.service.DeleteById(r.Context(), id); err != nil {
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}
	respond.NewResponse(w).DefaultMessage().Ok(nil)
}
