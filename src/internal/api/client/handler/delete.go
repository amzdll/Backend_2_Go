package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/nicklaw5/go-respond"
	"net/http"
)

// DeleteById
//
//	@Summary        Delete a client by ID
//	@Description    Deletes a client based on the provided ID
//	@Tags           Clients
//	@Accept         json
//	@Produce        json
//	@Param          id      path    string      true        "Client ID"
//	@Success        200     {object}    response.DefaultResponse      "Message successfully sent"
//	@Failure        400     {object}    response.ErrorResponse      "Bad Request"
//	@Router         /clients/{id} [delete]
func (h *Handler) DeleteById(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, err := uuid.Parse(idString)
	if err != nil {
		respond.NewResponse(w).DefaultMessage().BadRequest(nil)
		return
	}
	if err = h.service.DeleteById(r.Context(), id); err != nil {
		respond.NewResponse(w).DefaultMessage().NotFound(nil)
		return
	}
	respond.NewResponse(w).DefaultMessage().Ok(nil)
}
