package apifx

import (
	"context"
	clientHandler "github.com/amzdll/backend_2_go/src/internal/api/client/handler"
	clientRepository "github.com/amzdll/backend_2_go/src/internal/db/client/repository"
	clientService "github.com/amzdll/backend_2_go/src/internal/service/client"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
	"log"
	"net/http"
)

func Module() fx.Option {
	return fx.Module(
		"server",

		// Clients
		fx.Provide(
			fx.Annotate(clientRepository.New, fx.As(new(clientService.Repository))),
			fx.Annotate(clientService.New, fx.As(new(clientHandler.Service))),
			AsRoute(clientHandler.New),
		),

		// Common
		fx.Provide(
			NewConfig,
			NewValidator,
			fx.Annotate(MountHandlers, fx.ParamTags(`group:"routes"`)),
		),
		fx.Invoke(StartServer),
	)
}

func StartServer(lc fx.Lifecycle, router *chi.Mux, config *Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := http.ListenAndServe(config.Port, router)
			if err != nil {
				log.Fatalf("kekw: %v", err)
			}
			return nil
		},
		OnStop: func(context context.Context) error {
			return nil
		},
	})
}

type Route interface {
	Routes() *chi.Mux
}

func AsRoute(f interface{}) interface{} {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

func MountHandlers(routers []Route) *chi.Mux {
	mainRouter := chi.NewRouter()
	for _, router := range routers {
		mainRouter.Mount("/", router.Routes())
	}
	return mainRouter
}

func NewValidator() *validator.Validate {
	return validator.New()
}
