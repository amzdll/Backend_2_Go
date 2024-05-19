package app

import (
	"github.com/amzdll/backend_2_go/src/internal/app/apifx"
	"github.com/amzdll/backend_2_go/src/internal/app/dbfx"
	"go.uber.org/fx"
)

func CreateApp() *fx.App {
	return fx.New(
		fx.Options(
			dbfx.Module(),
			apifx.Module(),
		),
		fx.Provide(
			NewConfig,
		),
	)
}
