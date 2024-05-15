package service

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
)

func (cs ClientService) GetAll(ctx context.Context, listParams model.Pagination) ([]model.Client, error) {
	return cs.repository.GetAll(ctx, listParams)
}

func (cs ClientService) GetByNameSurname(
	ctx context.Context, clientInfo model.ClientInfo) ([]model.Client, error) {
	return cs.repository.GetByNameSurname(ctx, clientInfo)
}
