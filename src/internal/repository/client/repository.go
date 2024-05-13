package repository

import (
	"context"
	"fmt"
	"github.com/amzdll/backend_2_go/src/internal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientRepository struct {
	tableName string
	db        *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *ClientRepository {
	return &ClientRepository{tableName: "client", db: db}
}

func (repository ClientRepository) Create(ctx context.Context, client model.ClientInfo) error {
	query := fmt.Sprintf("INSERT INTO %s "+
		"(client_name, client_surname, birthday, gender, registration_date, address_id) "+
		"VALUES ($1, $2, $3, $4, $5, $6)", repository.tableName)

	_, err := repository.db.Exec(ctx, query,
		client.ClientName, client.ClientSurname, client.Birthday, client.Gender, client.RegistrationDate, client.AddressId)

	return err
}

func (repository ClientRepository) GetAll(ctx context.Context, clientListParams model.ClientListParams) ([]model.Client, error) {
	var clients []model.Client
	query := `select * from client`

	if err := pgxscan.Select(ctx, repository.db, &clients, query); err != nil {
		return nil, err
	}
	return clients, nil
}
