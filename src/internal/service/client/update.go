package service

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
)

func (cs ClientService) UpdateClient(ctx context.Context, data model.Client) error {
	return cs.repository.UpdateAddress(ctx, data)
}
