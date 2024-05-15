package handler

import (
	"github.com/amzdll/backend_2_go/src/internal/model"
	"github.com/nicklaw5/go-respond"
	"net/http"
)

func (h *Handler) GetByNameSurname(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	surname := r.URL.Query().Get("surname")
	if name == "" || surname == "" {
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}

	clients, err := h.service.GetByNameSurname(
		r.Context(), model.ClientInfo{ClientName: name, ClientSurname: surname},
	)

	if err != nil {
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}

	respond.NewResponse(w).Ok(clients)
}
