package service

import (
	"context"
	"github.com/google/uuid"
)

func (service ClientService) DeleteById(ctx context.Context, id uuid.UUID) error {
	return service.DeleteById(ctx, id)
}
