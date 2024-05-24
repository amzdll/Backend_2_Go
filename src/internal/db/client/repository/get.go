package repository

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
	"github.com/amzdll/backend_2_go/src/pkg/pgrepository"
)

func (r ClientRepository) GetAll(ctx context.Context, pagination model.Pagination) ([]model.Client, error) {
	return pgrepository.GetAll[model.Client](r.table, r.db, ctx, pagination)
}

func (r ClientRepository) GetByNameSurname(
	ctx context.Context, clientInfo model.ClientInfo,
) ([]model.Client, error) {
	//query := `select * from client where client_name = $1 and client_surname = $2`
	//err := pgxscan.Select(
	//	ctx, r.db, &clients, query, clientInfo.ClientName, clientInfo.ClientSurname)
	//if err != nil {
	//	return nil, err
	//}
	clients, err := pgrepository.GetWithFilters[model.Client](
		r.table,
		r.db,
		ctx,
		[]string{"client_name", "client_surname"}, clientInfo)

	if err != nil {
		return nil, err
	}

	return clients, nil
}
