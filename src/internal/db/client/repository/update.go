package repository

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
)

func (r ClientRepository) UpdateAddress(ctx context.Context, data model.Client) error {
	const query = `update client set address_id = $1 where id = $2`
	_, err := r.db.Exec(ctx, query, data.AddressId, data.Id)
	return err
}
