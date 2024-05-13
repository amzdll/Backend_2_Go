package main

import (
	"context"
	client "github.com/amzdll/backend_2_go/src/internal/api/client/handler"
	clientRepository "github.com/amzdll/backend_2_go/src/internal/repository/client"
	clientService "github.com/amzdll/backend_2_go/src/internal/service/client"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, os.Getenv("PG_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	validate := validator.New()

	app := chi.NewRouter()
	clientRepo := clientRepository.NewRepository(pool)
	clientServ := clientService.NewClientService(clientRepo)
	clientHandler := client.NewHandler(clientServ, validate)
	app.Mount("/", clientHandler.Routes())
	http.ListenAndServe(":8000", app)

	defer pool.Close()
}
