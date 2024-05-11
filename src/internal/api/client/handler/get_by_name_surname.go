package client

import (
	"net/http"
)

func (h *Handler) GetByNameSurname(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
