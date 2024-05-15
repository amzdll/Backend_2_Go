package repository

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
)

func (repository ClientRepository) Create(ctx context.Context, client model.ClientInfo) error {
	query := `insert into client (client_name, client_surname, birthday, gender, address_id) values ($1, $2, $3, $4, $5)`
	_, err := repository.db.Exec(
		ctx, query, client.ClientName, client.ClientSurname, client.Birthday, client.Gender, client.AddressId,
	)
	return err
}
