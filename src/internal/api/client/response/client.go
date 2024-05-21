package response

import (
	"github.com/google/uuid"
	"time"
)

type ClientResponse struct {
	Id               uuid.UUID `json:"id"`
	ClientName       string    `json:"client_name"`
	ClientSurname    string    `json:"client_surname"`
	Birthday         time.Time `json:"birthday"`
	Gender           string    `json:"gender"`
	RegistrationDate time.Time `json:"registration_date"`
	AddressId        uuid.UUID `json:"address_id"`
}

type ClientListResponse struct {
	Clients []ClientResponse
}
