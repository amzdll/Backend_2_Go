package service

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
	"github.com/google/uuid"
)

type ClientRepository interface {
	Create(ctx context.Context, data model.ClientInfo) error
	GetByNameSurname(ctx context.Context, data model.ClientInfo) ([]model.Client, error)
	GetAll(ctx context.Context, pagination model.Pagination) ([]model.Client, error)
	UpdateAddress(ctx context.Context, data model.Client) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}

type ClientService struct {
	repository ClientRepository
}

func NewClientService(cr ClientRepository) *ClientService {
	return &ClientService{repository: cr}
}
