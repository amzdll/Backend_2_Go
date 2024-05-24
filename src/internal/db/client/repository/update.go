package repository

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
)

func (r ClientRepository) UpdateAddress(ctx context.Context, data model.Client) error {
	_, err := r.db.Exec(ctx, queryUpdate, data.AddressId, data.Id)
	return err
}
