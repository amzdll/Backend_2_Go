package app

import (
	"github.com/amzdll/backend_2_go/src/internal/app/api"
	"github.com/amzdll/backend_2_go/src/internal/app/dbfx"
	"github.com/amzdll/backend_2_go/src/internal/config"
	"go.uber.org/fx"
)

func CreateApp() *fx.App {
	return fx.New(
		fx.Options(
			dbfx.Module(),
			api.Module(),
		),
		fx.Provide(
			config.NewConfig,
		),
	)
}
