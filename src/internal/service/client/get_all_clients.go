package client

import (
	"net/http"
)

func (h *Handler) GetAllClients(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello, World!"))
}
