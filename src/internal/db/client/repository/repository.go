package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientRepository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *ClientRepository {
	return &ClientRepository{db: db}
}
