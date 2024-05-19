package dbfx

import (
	"context"
	"fmt"
	"github.com/amzdll/backend_2_go/src/internal/config/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"db",
		fx.Provide(
			db.NewConfig,
			NewPool,
		),
	)
}

func NewPool(config *db.Config) (*pgxpool.Pool, error) {
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
