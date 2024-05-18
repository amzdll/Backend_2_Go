package http

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"log"
	"net/http"
)

func NewServerModule() fx.Option {
	return fx.Module(
		"server",
		fx.Provide(
			NewConfig,
			NewRouter,
		),
		fx.Invoke(
			RegisterRoutes,
			StartServer,
		),
	)
}

func NewRouter() chi.Router {
	return chi.NewRouter()
}

func RegisterRoutes(router chi.Router, handlers []Handler) {
	for _, h := range handlers {
		h.Register(router)
	}
}

func StartServer(lc fx.Lifecycle, router chi.Router, cfg *Config) {
	server := &http.Server{
		Addr:    cfg.Host + cfg.Port,
		Handler: router,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("HTTP server ListenAndServe: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down HTTP server...")
			return server.Shutdown(ctx)
		},
	})
}

type Handler interface {
	Register(router chi.Router)
}
