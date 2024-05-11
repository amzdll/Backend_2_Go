package client

import (
	"github.com/google/uuid"
	"time"
)

type CreateClientRequest struct {
	Id               uuid.UUID `json:"id,omitempty"`
	ClientName       string    `json:"client_name,omitempty"`
	ClientSurname    string    `json:"client_surname,omitempty"`
	Birthday         time.Time `json:"birthday"`
	Gender           string    `json:"gender,omitempty"`
	RegistrationDate time.Time `json:"registration_date"`
	AddressId        uuid.UUID `json:"address_id,omitempty"`
}
