package client

import (
	"encoding/json"
	"github.com/amzdll/backend_2_go/src/internal/converter"
	"github.com/go-chi/render"
	"net/http"

	clientReq "github.com/amzdll/backend_2_go/src/internal/api/client/request"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req clientReq.CreateClientRequest
	// need middleware
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, "Bad request")
		return
	}
	if err := h.validator.Struct(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, "Bad request")
		return
	}

	if err := h.service.Create(r.Context(), converter.ToClientInfoFromRequest(req)); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, "Bad request")
		return
	}

	w.WriteHeader(http.StatusCreated)
}
