package repository

import (
	"context"
	"github.com/amzdll/backend_2_go/src/pkg/pgrepository"
	"github.com/google/uuid"
)

func (r ClientRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	return pgrepository.Delete(r.table, r.db, ctx, id)
}
