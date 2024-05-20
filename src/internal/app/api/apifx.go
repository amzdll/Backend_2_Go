package api

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
	"log/slog"
	"net/http"
)

func Module() fx.Option {
	return fx.Module(
		"server",

		fx.Options(ClientModule()),

		fx.Provide(
			config.NewApiConfig,
			newValidator,
			fx.Annotate(mountHandlers, fx.ParamTags(`group:"routes"`)),
		),
		fx.Invoke(startServer),
	)
}

func startServer(lc fx.Lifecycle, router *chi.Mux, config *config.ApiConfig, l *slog.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			err := http.ListenAndServe(config.Port, router)
			if err != nil {
				l.Error("err", err)
			}
			return nil
		},
		OnStop: func(context context.Context) error {
			return nil
		},
	})
}

type route interface {
	Routes() *chi.Mux
}

func asRoute(f interface{}) interface{} {
	return fx.Annotate(
		f,
		fx.As(new(route)),
		fx.ResultTags(`group:"routes"`),
	)
}

func mountHandlers(routers []route) *chi.Mux {
	mainRouter := chi.NewRouter()

	mainRouter.Use(middleware.Logger)

	for _, router := range routers {
		mainRouter.Mount("/", router.Routes())
	}
	return mainRouter
}

func newValidator() *validator.Validate {
	return validator.New()
}
