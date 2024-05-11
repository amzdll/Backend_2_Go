package client

import (
	"net/http"
)

func (h *Handler) CreateClient(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello, World!"))
}
