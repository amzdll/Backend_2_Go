package api

import (
	"context"
	_ "github.com/amzdll/backend_2_go/src/api/shop"
	apiMiddleware "github.com/amzdll/backend_2_go/src/internal/api/middleware"
	"github.com/amzdll/backend_2_go/src/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/fx"
	"net/http"
)

func Module() fx.Option {
	return fx.Module(
		"server",
		fx.Options(ClientModule()),
		fx.Provide(
			fx.Annotate(setupMainRouter, fx.ParamTags(`group:"routes"`)),
		),
		fx.Provide(config.NewApiConfig),
		fx.Invoke(startServer),
	)
}

func startServer(lc fx.Lifecycle, router *chi.Mux, config *config.ApiConfig, l *zerolog.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			l.Info().Msg("Start server")
			err := http.ListenAndServe(config.Port, router)
			if err != nil {
				l.Error().Err(err).Msg("Cannot start server")
			}
			return nil
		},
		OnStop: func(context context.Context) error {
			return nil
		},
	})
}

func setupMainRouter(routers []route, cfg *config.ApiConfig, logger *zerolog.Logger) *chi.Mux {
	mainRouter := chi.NewRouter()
	lm := apiMiddleware.NewLogger(logger)
	mainRouter.Use(lm.Log)
	if cfg.Stage == config.EnvLocal || cfg.Stage == config.EnvDev {
		mainRouter.Get(
			"/swagger/*",
			httpSwagger.Handler(httpSwagger.URL("doc.json")),
		)
	}
	for _, router := range routers {
		mainRouter.Mount("/", router.Routes())
	}
	return mainRouter
}
