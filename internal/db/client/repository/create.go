package repository

import (
	"context"
	"github.com/amzdll/shop/internal/model"
)

func (r ClientRepository) Create(ctx context.Context, client model.ClientInfo) error {
	_, err := r.db.Exec(
		ctx, queryCreate, client.ClientName, client.ClientSurname, client.Birthday, client.Gender, client.AddressId,
	)
	return err
}
