package service

import (
	"context"
	"github.com/google/uuid"
)

func (s ClientService) DeleteById(ctx context.Context, id uuid.UUID) error {
	return s.repository.DeleteById(ctx, id)
}
