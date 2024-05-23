package db

import (
	"context"
	"fmt"
	"github.com/amzdll/backend_2_go/src/internal/app/logger"
	"github.com/amzdll/backend_2_go/src/internal/config/db"
	"github.com/jackc/pgx/v5/pgxpool"
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

func newPool(l *logger.Logger, config *db.Config) (*pgxpool.Pool, error) {
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
		l.Error("Failed to connect to database.", err)
		return nil, err
	}
	l.Info("Database connection pool has been opened.")
	return pool, nil
}

func closePool(lc fx.Lifecycle, pool *pgxpool.Pool, l *logger.Logger) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			pool.Close()
			l.Info("Database connection pool has been opened.")
			return nil
		},
	})
}
