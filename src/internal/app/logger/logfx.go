package logger

import (
	"github.com/amzdll/backend_2_go/src/internal/config"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"logger",
		fx.Provide(
			setupLogger,
		),
		fx.Provide(config.NewLogConfig),
	)
}
