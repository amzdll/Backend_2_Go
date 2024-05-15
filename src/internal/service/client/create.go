package service

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
)

func (cs ClientService) Create(ctx context.Context, clientInfo model.ClientInfo) error {
	return cs.repository.Create(ctx, clientInfo)
}
