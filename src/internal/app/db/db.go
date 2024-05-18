package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
	"log"
)

func NewDBModule() fx.Option {
	return fx.Module(
		"db",
		fx.Provide(
			NewConfig,
			NewPool,
		),
		fx.Invoke(registerHooks),
	)
}

func registerHooks(lc fx.Lifecycle, pool *pgxpool.Pool) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Database connection pool is starting...")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Closing database connection pool...")
			pool.Close()
			return nil
		},
	})
}

func NewPool(config *Config) (*pgxpool.Pool, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, config.Dsn)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
