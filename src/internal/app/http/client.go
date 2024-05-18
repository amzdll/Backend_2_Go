package http

import (
	"github.com/amzdll/backend_2_go/src/internal/api/client/handler"
	"github.com/amzdll/backend_2_go/src/internal/db/client/repository"
	service "github.com/amzdll/backend_2_go/src/internal/service/client"
	"go.uber.org/fx"
)

func NewClientModule() fx.Option {
	return fx.Module(
		"client",
		fx.Provide(
			repository.New,
			service.New,
			handler.New,
			NewConfig,
			NewValidator,
		),
		fx.Invoke(),
	)
}
