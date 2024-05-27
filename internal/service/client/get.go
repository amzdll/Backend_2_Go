package service

import (
	"context"
	"github.com/amzdll/shop/internal/model"
)

func (s ClientService) GetAll(ctx context.Context, listParams model.Pagination) ([]model.Client, error) {
	return s.repository.GetAll(ctx, listParams)
}

func (s ClientService) GetByNameSurname(
	ctx context.Context, clientInfo model.ClientInfo) ([]model.Client, error) {
	return s.repository.GetByNameSurname(ctx, clientInfo)
}
