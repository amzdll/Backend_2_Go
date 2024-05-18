package app

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/app/db"
	"github.com/amzdll/backend_2_go/src/internal/app/http"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"log"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Options(
			db.NewDBModule(),
			http.NewClientModule(),
		),
		fx.Provide(
			newRouter,
			NewConfig,
		),
		fx.Invoke(
			startServer,
			lifecycleHooks,
		),
	)
}

func newRouter() chi.Router {
	return chi.NewRouter()
}

//func startServer(lc fx.Lifecycle, router chi.Router, cfg *http.Config) {
//	server := &http.Server{
//		Addr:    cfg.Host + ":" + cfg.Port,
//		Handler: router,
//	}
//
//	lc.Append(fx.Hook{
//		OnStart: func(ctx context.Context) error {
//			go func() {
//				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
//					log.Fatalf("HTTP server ListenAndServe: %v", err)
//				}
//			}()
//			return nil
//		},
//		OnStop: func(ctx context.Context) error {
//			log.Println("Shutting down HTTP server...")
//			return server.Shutdown(ctx)
//		},
//	})
//}

func lifecycleHooks(lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Starting application...")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping application...")
			return nil
		},
	})
}
