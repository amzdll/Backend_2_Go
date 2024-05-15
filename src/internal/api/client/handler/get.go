package handler

import (
	"github.com/amzdll/backend_2_go/src/internal/api/client/converter"
	"github.com/amzdll/backend_2_go/src/internal/api/client/response"
	commonConverter "github.com/amzdll/backend_2_go/src/internal/api/common/converter"
	"github.com/amzdll/backend_2_go/src/internal/model"
	"github.com/nicklaw5/go-respond"
	"net/http"
)

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	pagination, err := commonConverter.RequestToPaginationModel(r)
	if err != nil {
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}
	clients, err := h.service.GetAll(r.Context(), pagination)
	if err != nil {
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}
	res := make([]response.ClientResponse, len(clients))
	for i, client := range clients {
		res[i] = converter.ModelToClientResponse(client)
	}
	respond.NewResponse(w).DefaultMessage().Ok(res)
}

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
