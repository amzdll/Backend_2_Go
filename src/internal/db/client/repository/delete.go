package repository

import (
	"context"
	"github.com/google/uuid"
)

func (repository ClientRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	const query = `delete from client where id = $1`
	_, err := repository.db.Exec(ctx, query, id)
	return err
}
