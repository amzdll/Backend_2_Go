package api

import (
	"context"
	"fmt"
	_ "github.com/amzdll/backend_2_go/src/api/shop"
	apiMiddleware "github.com/amzdll/backend_2_go/src/internal/api/middleware"
	"github.com/amzdll/backend_2_go/src/internal/app/logger"
	"github.com/amzdll/backend_2_go/src/internal/config"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/fx"
	"net/http"
)

func Module() fx.Option {
	return fx.Module(
		"server",
		fx.Options(ClientModule()),
		fx.Provide(fx.Annotate(setupMainRouter, fx.ParamTags(`group:"routes"`))),
		fx.Provide(config.NewApiConfig),
		fx.Invoke(startServer),
	)
}

func startServer(lc fx.Lifecycle, router *chi.Mux, config *config.ApiConfig, l *logger.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			l.Info("Server has started successfully.")
			err := http.ListenAndServe(config.Port, router)
			if err != nil {
				l.Error("Cannot start server", err)
			}
			return nil
		},
		OnStop: func(context context.Context) error {
			return nil
		},
	})
}

func setupMainRouter(routers []route, cfg *config.ApiConfig, logger *logger.Logger) *chi.Mux {
	mainRouter := chi.NewRouter()
	l := apiMiddleware.NewLogger(logger)
	mainRouter.Use(l.Logger)
	if cfg.Stage == config.EnvLocal || cfg.Stage == config.EnvDev {
		logger.Info(
			fmt.Sprintf(
				"Swagger documentation is available at http://%s%s/swagger/index.html",
				cfg.Host, cfg.Port,
			),
		)
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
