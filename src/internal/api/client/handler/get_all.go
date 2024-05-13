package handler

import (
	"github.com/amzdll/backend_2_go/src/internal/api/client/converter"
	"github.com/amzdll/backend_2_go/src/internal/api/client/request"
	"github.com/amzdll/backend_2_go/src/internal/api/client/response"
	"github.com/go-chi/render"
	"github.com/nicklaw5/go-respond"
	"net/http"
)

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	var req request.GetAllRequest

	if err := render.Decode(r, &req); err != nil {
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}
	clients, err := h.service.GetAll(r.Context(), converter.ToClientListParamsFromRequest(req))
	if err != nil {
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}

	res := make([]response.ClientResponse, len(clients))
	for i, client := range clients {
		res[i] = converter.ToClientResponseFromModel(client)
	}
	respond.NewResponse(w).DefaultMessage().Ok(res)
}
