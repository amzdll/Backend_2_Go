package api

import (
	clientHandler "github.com/amzdll/backend_2_go/src/internal/api/client/handler"
	clientRepository "github.com/amzdll/backend_2_go/src/internal/db/client/repository"
	clientService "github.com/amzdll/backend_2_go/src/internal/service/client"

	"go.uber.org/fx"
)

func ClientModule() fx.Option {
	return fx.Module(
		"server",

		fx.Provide(
			fx.Annotate(clientRepository.New, fx.As(new(clientService.Repository))),
			fx.Annotate(clientService.New, fx.As(new(clientHandler.Service))),
			AsRoute(clientHandler.New),
		))
}
