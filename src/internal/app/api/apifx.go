package api

import (
	"context"
	_ "github.com/amzdll/backend_2_go/src/api/shop"
	"github.com/amzdll/backend_2_go/src/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	httpSwagger "github.com/swaggo/http-swagger"
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
			fx.Annotate(setupMainRouter, fx.ParamTags(`group:"routes"`)),
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

func setupMainRouter(routers []route, cfg *config.ApiConfig) *chi.Mux {
	mainRouter := chi.NewRouter()
	//mainRouter.Use(cors.Handler(cors.Options{
	//	AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8000"},
	//	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	//	ExposedHeaders:   []string{"Link"},
	//	AllowCredentials: true,
	//	MaxAge:           300,
	//}))
	if cfg.Stage == config.EnvLocal {
		mainRouter.Use(middleware.Logger)

		mainRouter.Get(
			"/swagger/*",
			httpSwagger.Handler(
				httpSwagger.URL("doc.json"),
			))
	}
	for _, router := range routers {
		mainRouter.Mount("/", router.Routes())
	}
	return mainRouter
}

func newValidator() *validator.Validate {
	return validator.New()
}
