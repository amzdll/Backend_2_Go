package handler

import (
	"github.com/amzdll/shop/internal/api/client/response"
	commonConverter "github.com/amzdll/shop/internal/api/common/converter"
	"github.com/amzdll/shop/internal/model"
	"github.com/nicklaw5/go-respond"
	"net/http"
)

// GetAll
//
// @Summary Get all clients
// @Description Retrieves all clients
// @Tags Clients
// @Accept json
// @Produce json
// @Success 200 {array} response.ClientResponse "Clients successfully get"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Router /clients [get]
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
		res[i].From(client)
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
