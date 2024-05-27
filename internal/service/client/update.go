package service

import (
	"context"
	"github.com/amzdll/shop/internal/model"
)

func (s ClientService) UpdateClient(ctx context.Context, data model.Client) error {
	return s.repository.UpdateAddress(ctx, data)
}
