package app

import (
	"github.com/amzdll/backend_2_go/src/internal/app/api"
	"github.com/amzdll/backend_2_go/src/internal/app/db"
	"github.com/amzdll/backend_2_go/src/internal/app/logger"
	"github.com/amzdll/backend_2_go/src/internal/config"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log/slog"
)

func CreateApp() *fx.App {
	return fx.New(
		fx.Options(
			logger.Module(),
			db.Module(),
			api.Module(),
		),
		fx.Provide(
			config.NewConfig,
		),
		fx.WithLogger(func(log *slog.Logger) fxevent.Logger {
			return &fxevent.SlogLogger{Logger: log}
		}),
	)
}
