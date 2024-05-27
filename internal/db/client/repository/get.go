package repository

import (
	"context"
	"github.com/amzdll/shop/internal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r ClientRepository) GetAll(ctx context.Context, pagination model.Pagination) ([]model.Client, error) {
	var clients []model.Client
	if err := pgxscan.Select(ctx, r.db, &clients, queryGetAll, pagination.Limit, pagination.Offset); err != nil {
		return nil, err
	}
	return clients, nil
}

func (r ClientRepository) GetByNameSurname(
	ctx context.Context, clientInfo model.ClientInfo,
) ([]model.Client, error) {
	var clients []model.Client
	err := pgxscan.Select(
		ctx, r.db, &clients, queryGet, clientInfo.ClientName, clientInfo.ClientSurname,
	)
	if err != nil {
		return nil, err
	}
	return clients, nil
}
