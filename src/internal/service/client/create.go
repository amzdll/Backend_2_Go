package service

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
)

func (s ClientService) Create(ctx context.Context, clientInfo model.ClientInfo) error {
	return s.repository.Create(ctx, clientInfo)
}
