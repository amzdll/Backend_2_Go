package repository

import (
	"context"
	"github.com/amzdll/shop/internal/model"
)

func (r ClientRepository) UpdateAddress(ctx context.Context, data model.Client) error {
	_, err := r.db.Exec(ctx, queryUpdate, data.AddressId, data.Id)
	return err
}
