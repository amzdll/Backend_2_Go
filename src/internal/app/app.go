package app

import (
	"github.com/amzdll/backend_2_go/src/internal/app/apifx"
	"github.com/amzdll/backend_2_go/src/internal/app/dbfx"
	"github.com/amzdll/backend_2_go/src/internal/config"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"go.uber.org/fx"
)

// CreateApp start of integration Zap Logger
func CreateApp() *fx.App {
	return fx.New(
		fx.Options(
			dbfx.PgModule(),
			apifx.Module(),
		),
		fx.Provide(
			config.NewConfig,
		),
		fx.Provide(
			zap.NewProduction,
		),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)
}
