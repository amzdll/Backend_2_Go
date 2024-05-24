package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (r ClientRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.Exec(ctx, queryDelete, id)
	if err == nil && result.RowsAffected() == 0 {
		err = fmt.Errorf("not found")
	}
	return err
}
