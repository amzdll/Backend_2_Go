package pgcrud

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

const queryCreate = `insert into %s (%s) values (%s)`
const queryGet = `select * from %s where %s and %s`
const queryGetAll = `select * from %s limit $1 offset $2`
const queryUpdate = `update %s set address_id = $1 where id = $2`
const queryDelete = `delete from %s where %s`

func Create(
	ctx context.Context, db *pgxpool.Pool, table string, fields string,
) error {
	return nil
}

func Get(
	ctx context.Context, db *pgxpool.Pool, table string, fields string,
) error {
	return nil
}

func GetAll(
	ctx context.Context, db *pgxpool.Pool, table string, fields string,
) error {
	return nil
}

func Update(
	ctx context.Context, db *pgxpool.Pool, table string, fields string,
) error {
	return nil
}

func Delete(
	ctx context.Context, db *pgxpool.Pool, table string, fields string,
) error {
	return nil
}
