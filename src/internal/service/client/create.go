package service

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
)

func (service ClientService) Create(ctx context.Context, clientInfo model.ClientInfo) error {
	return service.repository.Create(ctx, clientInfo)
}
