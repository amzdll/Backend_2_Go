package apifx

import (
	"context"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"log"
	"net/http"
)

func Module() fx.Option {
	return fx.Module(
		"server",

		fx.Options(
			ClientModule(),
		),

		fx.Provide(
			NewConfig,
			fx.Annotate(
				MountHandlers,
				fx.ParamTags(`group:"routes"`),
			),
		),
		fx.Invoke(StartServer),
	)
}

func StartServer(lc fx.Lifecycle, router *chi.Mux, config *Config) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			err := http.ListenAndServe(config.Port, router)
			if err != nil {
				log.Fatalf("Ошибка проверки соединения с базой данных: %v", err)
			}
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
