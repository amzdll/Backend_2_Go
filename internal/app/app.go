package app

import (
	"github.com/amzdll/shop/internal/app/api"
	"github.com/amzdll/shop/internal/app/db"
	"github.com/amzdll/shop/internal/app/logger"
	"github.com/amzdll/shop/internal/app/validator"
	"github.com/amzdll/shop/internal/config"
	"go.uber.org/fx"
)

func CreateApp() *fx.App {
	return fx.New(
		fx.Options(
			logger.Module(),
			db.Module(),
			validator.Module(),
			api.Module(),
		),
		fx.Provide(
			config.NewConfig,
		),
		fx.NopLogger,
	)
}
