package service

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
)

func (service ClientService) GetAll(ctx context.Context, listParams model.Pagination) ([]model.Client, error) {
	return service.repository.GetAll(ctx, listParams)
}
