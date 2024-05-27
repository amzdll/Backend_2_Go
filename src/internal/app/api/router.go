package api

import (
	"fmt"
	apiMiddleware "github.com/amzdll/backend_2_go/src/internal/api/middleware"
	"github.com/amzdll/backend_2_go/src/internal/app/logger"
	"github.com/amzdll/backend_2_go/src/internal/config"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/fx"
)

func setupMainRouter(routers []route, cfg *config.ApiConfig, logger *logger.Logger) *chi.Mux {
	mainRouter := chi.NewRouter()
	l := apiMiddleware.NewLogger(logger)
	mainRouter.Use(l.Logger)
	if cfg.Stage == config.EnvLocal || cfg.Stage == config.EnvDev {
		logger.Info(
			fmt.Sprintf(
				"Swagger documentation is available at http://%s%s/swagger/index.html",
				cfg.Host, cfg.Port,
			),
		)
		mainRouter.Get(
			"/swagger/*",
			httpSwagger.Handler(httpSwagger.URL("doc.json")),
		)
	}
	for _, router := range routers {
		mainRouter.Mount("/", router.Routes())
	}
	return mainRouter
}

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
