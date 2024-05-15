package repository

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (repository ClientRepository) GetByNameSurname(
	ctx context.Context, clientInfo model.ClientInfo,
) ([]model.Client, error) {
	var clients []model.Client

	query := `select * from client where client_name = $1 and client_surname = $2`
	err := pgxscan.Select(
		ctx, repository.db, &clients, query, clientInfo.ClientName, clientInfo.ClientSurname)
	if err != nil {
		return nil, err
	}

	return clients, nil
}
