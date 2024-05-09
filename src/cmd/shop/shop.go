package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	http.ListenAndServe(":3000", r)
}
