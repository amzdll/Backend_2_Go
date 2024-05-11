package main

import (
	"github.com/amzdll/backend_2_go/src/internal/api/client/handler"
	services "github.com/amzdll/backend_2_go/src/internal/service"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {

	r := chi.NewRouter()
	s := services.NewClientService()
	clientHandler := client.NewHandler(s)
	r.Mount("/", clientHandler.Routes())

	http.ListenAndServe(":8000", r)
}
