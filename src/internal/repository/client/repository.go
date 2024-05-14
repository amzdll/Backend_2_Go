package repository

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *ClientRepository {
	return &ClientRepository{db: db}
}

func (repository ClientRepository) Create(ctx context.Context, client model.ClientInfo) error {
	query := `insert into client 
    					(client_name, client_surname, birthday, gender, address_id)
						values ($1, $2, $3, $4, $5)`

	_, err := repository.db.Exec(ctx, query,
		client.ClientName, client.ClientSurname, client.Birthday, client.Gender, client.AddressId)
	return err
}

func (repository ClientRepository) GetByNameSurname(
	ctx context.Context, clientInfo model.ClientInfo,
) ([]model.Client, error) {
	query := `select * from client where client_name = $1 and client_surname = $2`
	var clients []model.Client

	err := pgxscan.Select(
		ctx, repository.db, &clients, query, clientInfo.ClientName, clientInfo.ClientSurname)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

func (repository ClientRepository) GetAll(ctx context.Context, pagination model.Pagination) ([]model.Client, error) {
	query := `select * from client limit $1 offset $2`
	var clients []model.Client

	if err := pgxscan.Select(ctx, repository.db, &clients, query, pagination.Limit, pagination.Offset); err != nil {
		return nil, err
	}

	return clients, nil
}

func (repository ClientRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := `delete from client where id = $1`
	_, err := repository.db.Exec(ctx, query, id)
	return err
}
