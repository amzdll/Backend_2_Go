package handler

import (
	"github.com/amzdll/backend_2_go/src/internal/api/client/converter"
	"github.com/amzdll/backend_2_go/src/internal/api/client/request"
	"github.com/go-chi/render"
	"github.com/nicklaw5/go-respond"
	"net/http"
)

func (h *Handler) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	var req request.UpdationRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		h.logger.Debug("Invalid request", "err", err)
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}
	if err := h.validator.Struct(req); err != nil {
		h.logger.Debug("Invalid parameters", "err", err)
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}
	if err := h.service.UpdateClient(r.Context(), converter.ToClientFromUpdateAddressRequest(req)); err != nil {
		h.logger.Debug("", "err", err)
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}
	respond.NewResponse(w).DefaultMessage().Ok(nil)
}
