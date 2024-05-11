package client

import (
	"github.com/google/uuid"
	"time"
)

type CreateClientRequest struct {
	Id               uuid.UUID
	ClientName       string
	ClientSurname    string
	Birthday         time.Time
	Gender           string
	RegistrationDate time.Time
	AddressId        uuid.UUID
}
