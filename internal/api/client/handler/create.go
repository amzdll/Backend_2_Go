package handler

import (
	"github.com/amzdll/shop/internal/api/client/request"
	"github.com/nicklaw5/go-respond"

	"github.com/go-chi/render"
	"net/http"
)

// Create
//
//	@Summary			Create new client
//	@Description	Create new client
//	@Tags			Clients
//	@Accept			json
//	@Produce		json
//	@Param			text	body		request.CreateRequest		true	"Request payload"
//	@Success		201		{object}	response.DefaultResponse	"Message successfully sent"
//	@Failure		400		{object}	response.ErrorResponse		"Bad Request"
//	@Router			/clients [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req request.CreateRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}
	if err := h.validator.Struct(req); err != nil {
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}
	if err := h.service.Create(r.Context(), req.ToClientInfo()); err != nil {
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}
	respond.NewResponse(w).DefaultMessage().Created(nil)
}
