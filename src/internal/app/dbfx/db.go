package dbfx

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"db",
		fx.Provide(
			NewConfig,
			NewPool,
		),
	)
}

func NewPool(config *Config) (*pgxpool.Pool, error) {
	ctx := context.Background()
	dsn := fmt.Sprintf(
		config.DsnTemplate,
		config.Host, config.Port, config.Name,
		config.User, config.Password, config.SslMode,
	)
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
