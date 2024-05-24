package repository

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
	"github.com/amzdll/backend_2_go/src/pkg/pgrepository"
)

func (r ClientRepository) Create(ctx context.Context, client model.ClientInfo) error {
	return pgrepository.Create(r.table, r.db, ctx, client)
}
