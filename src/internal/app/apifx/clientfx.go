package apifx

import (
	"github.com/amzdll/backend_2_go/src/internal/api/client/handler"
	"github.com/amzdll/backend_2_go/src/internal/db/client/repository"
	service "github.com/amzdll/backend_2_go/src/internal/service/client"
	"go.uber.org/fx"
)

func ClientModule() fx.Option {
	return fx.Module(
		"client",

		fx.Provide(
			NewValidator,
		),

		fx.Provide(
			fx.Annotate(
				repository.New,
				fx.As(new(service.ClientRepository)),
			),

			fx.Annotate(
				service.New,
				fx.As(new(handler.Service)),
			),

			AsRoute(handler.New),
		),
	)
}
