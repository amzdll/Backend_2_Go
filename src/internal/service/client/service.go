package service

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
)

type ClientRepository interface {
	Create(ctx context.Context, clientInfo model.ClientInfo) error
	//GetByNameSurname() (*model.Client, error)
	GetAll(ctx context.Context, clientListParams model.ClientListParams) ([]model.Client, error)
	//UpdateAddress() error
	//DeleteById() error
}

type ClientService struct {
	repository ClientRepository
}

func NewClientService(repository ClientRepository) *ClientService {
	return &ClientService{repository: repository}
}
