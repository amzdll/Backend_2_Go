package logger

import (
	"github.com/amzdll/shop/internal/config"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"logger",
		fx.Provide(NewLogger),
		fx.Provide(config.NewLogConfig),
	)
}
