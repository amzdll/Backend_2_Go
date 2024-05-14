package service

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
	"github.com/google/uuid"
)

type ClientRepository interface {
	Create(ctx context.Context, clientInfo model.ClientInfo) error
	GetByNameSurname(ctx context.Context, clientInfo model.ClientInfo) ([]model.Client, error)
	GetAll(ctx context.Context, pagination model.Pagination) ([]model.Client, error)
	//UpdateAddress() error
	DeleteById(ctx context.Context, id uuid.UUID) error
}

type ClientService struct {
	repository ClientRepository
}

func NewClientService(repository ClientRepository) *ClientService {
	return &ClientService{repository: repository}
}
