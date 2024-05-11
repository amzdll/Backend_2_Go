package main

import (
	client "github.com/amzdll/backend_2_go/src/internal/api/client/handler"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	clientHandler := client.NewHandler()

	r := chi.NewRouter()
	r.Mount("/", clientHandler.Routes())

	http.ListenAndServe(":8000", r)
}
