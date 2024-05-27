package service

import (
	"context"
	"github.com/amzdll/shop/internal/model"
)

func (s ClientService) Create(ctx context.Context, clientInfo model.ClientInfo) error {
	return s.repository.Create(ctx, clientInfo)
}
