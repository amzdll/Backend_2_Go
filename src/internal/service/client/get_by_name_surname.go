package service

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
)

func (service ClientService) GetByNameSurname(
	ctx context.Context, clientInfo model.ClientInfo) ([]model.Client, error) {
	return service.repository.GetByNameSurname(ctx, clientInfo)
}
