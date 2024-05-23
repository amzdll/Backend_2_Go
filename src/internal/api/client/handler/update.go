package handler

import (
	"github.com/amzdll/backend_2_go/src/internal/api/client/request"
	"github.com/go-chi/render"
	"github.com/nicklaw5/go-respond"
	"net/http"
)

// UpdateAddress
//
//	@Summary        Update client address
//	@Description    Updates a client's address
//	@Tags           Clients
//	@Accept         json
//	@Produce        json
//	@Param          body    body    request.UpdateRequest     true    "Update Client Address Request"
//	@Success        200     {object}    response.DefaultResponse     "Message successfully sent"
//	@Failure        400     {object}    response.ErrorResponse      "bad request"
//	@Router         /clients [put]
func (h *Handler) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	var req request.UpdateRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}
	if err := h.validator.Struct(req); err != nil {
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}
	if err := h.service.UpdateClient(r.Context(), req.ToClient()); err != nil {
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}
	respond.NewResponse(w).DefaultMessage().Ok(nil)
}
