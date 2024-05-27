package api

import (
	"context"
	"fmt"
	_ "github.com/amzdll/shop/api"
	apiMiddleware "github.com/amzdll/shop/internal/api/middleware"
	"github.com/amzdll/shop/internal/app/logger"
	"github.com/amzdll/shop/internal/config"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/fx"
	"net"
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
			listener, err := net.Listen("tcp", config.Port)
			if err != nil {
				l.Error("Cannot start server", err)
				return err
			}
			l.Info("Server has started successfully.")
			if err = http.Serve(listener, router); err != nil {
				l.Error("Cannot start server", err)
				return err
			}
			l.Info("Server has started successfully.")
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
