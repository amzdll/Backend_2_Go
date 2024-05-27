package api

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type route interface {
	Routes() *chi.Mux
}

func asRoute(f interface{}) interface{} {
	return fx.Annotate(
		f,
		fx.As(new(route)),
		fx.ResultTags(`group:"routes"`),
	)
}
