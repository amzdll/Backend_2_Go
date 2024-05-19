package api

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
	"log"
	"net/http"
)

func Module() fx.Option {
	return fx.Module(
		"server",

		fx.Options(ClientModule()),

		fx.Provide(
			config.NewApiConfig,
			NewValidator,
			fx.Annotate(MountHandlers, fx.ParamTags(`group:"routes"`)),
		),
		fx.Invoke(StartServer),
	)
}

func StartServer(lc fx.Lifecycle, router *chi.Mux, config *config.ApiConfig) {
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
