package client

import (
	"net/http"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.service.ClientHello()))
}
