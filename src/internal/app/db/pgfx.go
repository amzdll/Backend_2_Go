package db

import (
	"context"
	"fmt"
	"github.com/amzdll/backend_2_go/src/internal/config/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

func PgModule() fx.Option {
	return fx.Module(
		"pg",
		fx.Provide(newPool),
		fx.Provide(db.NewConfig),
		fx.Invoke(closePool),
	)
}

func newPool(l *zerolog.Logger, config *db.Config) (*pgxpool.Pool, error) {
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
	if err = pool.Ping(ctx); err != nil {
		l.Error().Err(err).Msg("Failed to connect to database")
	}
	l.Info().Msg("Database connection pool is open")
	return pool, nil
}

func closePool(lc fx.Lifecycle, pool *pgxpool.Pool, l *zerolog.Logger) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			pool.Close()
			l.Info().Msg("Database connection pool closed")
			return nil
		},
	})
}
