package service

import (
	"context"
	"github.com/google/uuid"
)

func (cs ClientService) DeleteById(ctx context.Context, id uuid.UUID) error {
	return cs.repository.DeleteById(ctx, id)
}
